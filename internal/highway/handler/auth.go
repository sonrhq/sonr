package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/protocol"

	mdw "github.com/sonrhq/core/internal/highway/middleware"
	"github.com/sonrhq/core/internal/highway/types"
	authenticationpb "github.com/sonrhq/core/types/highway/authentication/v1"
	domainproxy "github.com/sonrhq/core/x/domain/client/proxy"
	identityproxy "github.com/sonrhq/core/x/identity/client/proxy"
	serviceproxy "github.com/sonrhq/core/x/service/client/proxy"
)

// AuthenticationAPI is the handler for the authentication service server.
type AuthenticationAPI struct {
	authenticationpb.AuthenticationServiceServer
}

// CurrentUser returns the current user's details.
//
// @Summary Current user's details
// @Description Returns the current user's details.
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string "JWT Token"
// @Failure 400 {object} map[string]string "Missing JWT"
// @Failure 401 {object} map[string]string "No JWT cookie found"
// @Router /currentUser [get]
func CurrentUser(c *gin.Context) {
	jwt, err := c.Cookie("sonr-jwt")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "status": "no jwt cookie found"})
		return
	}
	if jwt == "" {
		c.JSON(400, gin.H{
			"error": "missing jwt",
		})
		return
	}
	c.JSON(200, gin.H{
		"jwt": jwt,
	})
}

// RegisterEscrowIdentity returns the credential assertion options to start account login.
//
// @Summary Register escrow identity
// @Description Returns the credential assertion options to start account login.
// @Accept  json
// @Produce  json
// @Param amount query string true "Amount"
// @Param email query string true "Email"
// @Success 200 {object} map[string]interface{} "Success response"
// @Failure 500 {object} map[string]string "Error message"
// @Router /registerEscrowIdentity [get]
func RegisterEscrowIdentity(c *gin.Context) {
	origin := c.Query("amount")
	alias := c.Query("email")

	record, err := serviceproxy.GetServiceRecord(c.Request.Context(), origin)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "where": "GetServiceRecord"})
		return
	}
	assertionOpts, chal, _, err := mdw.IssueCredentialAssertionOptions(alias, record)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "where": "GetCredentialAssertionOptions"})
		return
	}
	c.JSON(200, gin.H{
		"assertion_options": assertionOpts,
		"challenge":         chal.String(),
		"origin":            origin,
		"alias":             alias,
	})
}

// RegisterControllerIdentity registers a credential for a given Unclaimed Wallet.
//
// @Summary Register controller identity
// @Description Registers a credential for a given Unclaimed Wallet.
// @Accept  json
// @Produce  json
// @Param origin path string true "Origin"
// @Param alias path string true "Alias"
// @Param attestation query string true "Attestation"
// @Param challenge query string true "Challenge"
// @Success 200 {object} map[string]interface{} "Success response"
// @Failure 404 {object} map[string]string "Error message"
// @Failure 412 {object} map[string]string "Error message"
// @Failure 400 {object} map[string]string "Error message"
// @Failure 401 {object} map[string]string "Error message"
// @Router /registerControllerIdentity/{origin}/{alias} [post]
func RegisterControllerIdentity(c *gin.Context) {
	origin := c.Param("origin")
	alias := c.Param("alias")
	attestionResp := c.Query("attestation")
	challenge := c.Query("challenge")
	// Get the service record from the origin
	record, err := serviceproxy.GetServiceRecord(c.Request.Context(), origin)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error(), "where": "GetServiceRecord"})
		return
	}
	credential, err := record.VerifyCreationChallenge(attestionResp, challenge)
	if err != nil && credential == nil {
		c.JSON(412, gin.H{"error": err.Error(), "where": "VerifyCreationChallenge"})
		return
	}
	cont, resp, err := mdw.PublishControllerAccount(alias, credential, origin)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "where": "PublishControllerAccount"})
		return
	}
	token, err := types.NewSessionJWTClaims(alias, cont.Account())
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error(), "where": "NewSessionJWTClaims"})
		return
	}
	c.JSON(200, gin.H{
		"tx_hash": resp.TxHash,
		"address": cont.Account().Address,
		"token":   token,
		"origin":  origin,
		"success": true,
	})
}

// SignInWithCredential verifies a credential for a given account and returns the JWT Keyshare.
//
// @Summary Sign in with credential
// @Description Verifies a credential for a given account and returns the JWT Keyshare.
// @Accept  json
// @Produce  json
// @Param origin path string true "Origin"
// @Param alias path string true "Alias"
// @Param assertion query string true "Assertion"
// @Success 200 {object} map[string]interface{} "Success response"
// @Failure 441 {object} map[string]string "Error message"
// @Failure 442 {object} map[string]string "Error message"
// @Failure 443 {object} map[string]string "Error message"
// @Failure 444 {object} map[string]string "Error message"
// @Failure 445 {object} map[string]string "Error message"
// @Router /signInWithCredential/{origin}/{alias} [post]
func SignInWithCredential(c *gin.Context) {
	origin := c.Param("origin")
	alias := c.Param("alias")
	assertionResp := c.Query("assertion")
	record, err := serviceproxy.GetServiceRecord(c.Request.Context(), origin)
	if err != nil {
		c.JSON(441, gin.H{"error": err.Error(), "where": "GetServiceRecord"})
		return
	}
	_, err = record.VerifyAssertionChallenge(assertionResp)
	if err != nil {
		c.JSON(442, gin.H{"error": err.Error(), "where": "VerifyCreationChallenge"})
		return
	}
	addr, err := domainproxy.GetEmailRecordCreator(c.Request.Context(), alias)
	if err != nil {
		c.JSON(443, gin.H{"error": err.Error(), "where": "GetEmailRecordCreator"})
		return
	}
	contAcc, err := identityproxy.GetControllerAccount(c.Request.Context(), addr)
	if err != nil {
		c.JSON(444, gin.H{"error": err.Error(), "where": "GetControllerAccount"})
		return
	}
	token, err := types.NewSessionJWTClaims(alias, contAcc)
	if err != nil {
		c.JSON(445, gin.H{"error": err.Error(), "where": "NewSessionJWTClaims"})
		return
	}
	c.JSON(200, gin.H{
		"token":   token,
		"origin":  origin,
		"address": contAcc.Address,
		"success": true,
	})
}

// SignInWithEmail registers a DIDDocument for a given email authorized jwt.
//
// @Summary Sign in with email
// @Description Registers a DIDDocument for a given email authorized jwt.
// @Accept  json
// @Produce  json
// @Param jwt query string true "JWT"
// @Success 200 {object} map[string]interface{} "Success response"
// @Failure 500 {object} map[string]string "Error message"
// @Router /signInWithEmail [post]
func SignInWithEmail(c *gin.Context) {
	// token := c.Query("jwt")
	// resp, err := sfs.Claims.ClaimWithEmail(token)
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": err.Error(), "where": "ClaimWithEmail"})
	// 	return
	// }
	// c.JSON(200, mdw.StoreAuthCookies(c, resp, ""))
	c.JSON(500, gin.H{
		// "jwt":     token,
		// "ucw_id":  ucw,
		// "address": ucw,
	})
}

// GetCredentialAttestationParams returns the credential creation options to start account registration.
//
// @Summary Get credential attestation parameters
// @Description Returns the credential creation options to start account registration.
// @Accept  json
// @Produce  json
// @Param origin path string true "Origin"
// @Param alias path string true "Alias"
// @Success 200 {object} map[string]interface{} "Success response"
// @Failure 500 {object} map[string]string "Error message"
// @Router /getCredentialAttestationParams/{origin}/{alias} [get]
func GetCredentialAttestationParams(c *gin.Context) {
	origin := c.Param("origin")
	alias := c.Param("alias")
	ok, err := domainproxy.CheckAliasAvailable(c.Request.Context(), alias)
	if err != nil && !ok {
		c.JSON(500, gin.H{"error": err.Error(), "where": "CheckAliasAvailable"})
		return
	}
	// Get the service record from the origin
	rec, err := serviceproxy.GetServiceRecord(c.Request.Context(), origin)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "where": "GetServiceRecord"})
		return
	}
	chal, err := protocol.CreateChallenge()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "where": "CreateChallenge"})
		return
	}
	creOpts, err := rec.GetCredentialCreationOptions(alias, chal)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "where": "GetCredentialCreationOptions"})
		return
	}
	c.JSON(200, gin.H{
		"origin":            origin,
		"alias":             alias,
		"attestion_options": creOpts,
		"challenge":         chal.String(),
	})
}

// GetCredentialAssertionParams returns the credential assertion options to start account login.
//
// @Summary Get credential assertion parameters
// @Description Returns the credential assertion options to start account login.
// @Accept  json
// @Produce  json
// @Param origin path string true "Origin"
// @Param alias path string true "Alias"
// @Success 200 {object} map[string]interface{} "Success response"
// @Failure 500 {object} map[string]string "Error message"
// @Router /getCredentialAssertionParams/{origin}/{alias} [get]
func GetCredentialAssertionParams(c *gin.Context) {
	origin := c.Param("origin")
	alias := c.Param("alias")
	record, err := serviceproxy.GetServiceRecord(c.Request.Context(), origin)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "where": "GetServiceRecord"})
		return
	}
	notok, err := domainproxy.CheckAliasUnavailable(c.Request.Context(), alias)
	if err != nil && notok {
		c.JSON(500, gin.H{"error": err.Error(), "where": "CheckAliasAvailable"})
		return
	}
	assertionOpts, chal, addr, err := mdw.IssueCredentialAssertionOptions(alias, record)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error(), "where": "GetCredentialAssertionOptions"})
		return
	}
	c.JSON(200, gin.H{
		"assertion_options": assertionOpts,
		"challenge":         chal.String(),
		"origin":            origin,
		"alias":             alias,
		"address":           addr,
	})
}

// GetEmailAssertionParams returns a JWT for the email controller. After it is confirmed, the user will claim one of their unclaimed Keyshares.
//
// @Summary Get email assertion parameters
// @Description Returns a JWT for the email controller. After it is confirmed, the user will claim one of their unclaimed Keyshares.
// @Accept  json
// @Produce  json
// @Param email query string true "Email"
// @Success 200 {object} map[string]interface{} "Success response"
// @Failure 500 {object} map[string]string "Error message"
// @Router /getEmailAssertionParams [get]
func GetEmailAssertionParams(c *gin.Context) {
	// email := c.Query("email")
	// ucw, err := sfs.Accounts.RandomUnclaimedWallet()
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": err.Error(), "where": "RandomUnclaimedWallet"})
	// 	return
	// }
	// token, err := sfs.Claims.IssueEmailClaims(email, ucw)
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": err.Error(), "where": "IssueEmailClaims"})
	// 	return
	// }
	c.JSON(500, gin.H{
		// "jwt":     token,
		// "ucw_id":  ucw,
		// "address": ucw,
	})
}
