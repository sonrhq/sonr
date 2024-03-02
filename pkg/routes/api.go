package routes

import (
	"github.com/labstack/echo/v4"

	api "github.com/sonrhq/sonr/pkg/handlers/api"
)

// RegisterCosmosAPI registers the Cosmos API routes
func RegisterCosmosAPI(e *echo.Echo) {
	e.GET("/balance/:address", api.BankAPI.GetAllBalances)
	e.GET("/balance/:address/:denom", api.BankAPI.GetBalance)
	e.GET("/balance/:address/:denom/spendable", api.BankAPI.GetSpendableBalancesByDenom)
	e.GET("/balance/:address/spendable", api.BankAPI.GetSpendableBalances)
	e.GET("/block", api.CometAPI.GetLatestBlock)
	e.GET("/block/:height", api.CometAPI.GetBlockByHeight)
	e.GET("/constitution", api.GovAPI.GetConstitution)
	e.GET("/delegators/:delegatorAddr", api.StakingAPI.GetDelegatorDelegations)
	e.GET("/delegators/:delegatorAddr/validators", api.StakingAPI.GetDelegatorValidators)
	e.GET("/delegators/:delegatorAddr/validators/:validatorAddr", api.StakingAPI.GetDelegation)
	e.GET("/delegators/:delegatorAddr/validators/:validatorAddr/unbonding", api.StakingAPI.GetUnbondingDelegation)
	e.GET("/delegators/:delegatorAddr/validators/:srcValidatorAddr/redelegate/:dstValidatorAddr", api.StakingAPI.GetRedelegations)
	e.GET("/delegators/:delegatorAddr/unbonding", api.StakingAPI.GetDelegatorUnbondingDelegations)
	e.GET("/health", api.CometAPI.GetNodeInfo)
	e.GET("/history/{height}", api.StakingAPI.GetHistoricalInfo)
	e.GET("/proposals", api.GovAPI.GetProposals)
	e.GET("/proposals/:proposalId", api.GovAPI.GetProposal)
	e.GET("/proposals/:proposalId/deposits", api.GovAPI.GetDeposits)
	e.GET("/proposals/:proposalId/deposits/:depositor", api.GovAPI.GetDeposit)
	e.GET("/proposals/:proposalId/tally", api.GovAPI.GetTally)
	e.GET("/proposals/:proposalId/votes", api.GovAPI.GetVotes)
	e.GET("/proposals/:proposalId/votes/:voter", api.GovAPI.GetVote)
	e.GET("/staking", api.StakingAPI.GetValidators)
	e.GET("/staking/:validatorAddr", api.StakingAPI.GetValidator)
	e.GET("/staking/:validatorAddr/delegations", api.StakingAPI.GetValidatorDelegations)
	e.GET("/supply", api.BankAPI.GetTotalSupply)
	e.GET("/supply/:denom", api.BankAPI.GetSupplyOf)
	e.GET("/syncing", api.CometAPI.GetSyncing)
	e.GET("/validators", api.CometAPI.GetLatestValidatorSet)
	e.GET("/validators/:height", api.CometAPI.GetValidatorSetByHeight)
}

// RegisterSonrAPI registers the Sonr API routes
func RegisterSonrAPI(e *echo.Echo) {
	e.GET("/check/identifier/:id", api.AuthAPI.CheckIdentifier)
	e.GET("/service/:origin", api.ServiceAPI.QueryOrigin)
	e.GET("/service/:origin/login/:username/start", api.ServiceAPI.StartLogin)
	e.POST("/service/:origin/login/:username/finish", api.ServiceAPI.FinishLogin)
	e.GET("/service/:origin/register/:username/start", api.ServiceAPI.StartRegistration)
	e.POST("/service/:origin/register/:username/finish", api.ServiceAPI.FinishRegistration)
	e.GET("/tx/:txHash", api.TxAPI.GetTx)
	e.GET("/tx/block/:height", api.TxAPI.GetBlockWithTxs)
	e.POST("/tx/broadcast", api.TxAPI.BroadcastTx)
	e.POST("/tx/simulate", api.TxAPI.SimulateTx)
}
