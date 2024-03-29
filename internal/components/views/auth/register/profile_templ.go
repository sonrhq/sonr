// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package register

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/didao-org/sonr/internal/components/views/auth"
)

func createProfileScreen() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<main class=\"grow\"><section class=\"relative\"><!-- Illustration --><div class=\"md:block absolute left-1/2 -translate-x-1/2 -mt-36 blur-2xl opacity-70 pointer-events-none -z-10\" aria-hidden=\"true\"><img src=\"https://cdn.sonr.build/images/auth-illustration.svg\" class=\"max-w-none\" width=\"1440\" height=\"450\" alt=\"Page Illustration\"></div><div class=\"relative max-w-6xl mx-auto px-4 sm:px-6\"><div class=\"pt-32 pb-12 md:pt-40 md:pb-20\"><!-- Page header --><div class=\"max-w-3xl mx-auto text-center pb-12\"><!-- Logo --><div class=\"mb-5\"><a class=\"inline-flex\" href=\"/\"><div class=\"relative flex items-center justify-center w-16 h-16 border border-transparent rounded-2xl shadow-2xl [background:linear-gradient(theme(colors.slate.900),_theme(colors.slate.900))_padding-box,_conic-gradient(theme(colors.slate.400),_theme(colors.slate.700)_25%,_theme(colors.slate.700)_75%,_theme(colors.slate.400)_100%)_border-box] before:absolute before:inset-0 before:bg-slate-800/30 before:rounded-2xl\"><img class=\"relative\" src=\"https://cdn.sonr.build/images/logo.svg\" width=\"42\" height=\"42\" alt=\"Sonr\"></div></a></div><!-- Page title --><h1 class=\"h2 bg-clip-text text-transparent bg-gradient-to-r from-slate-200/60 via-slate-200 to-slate-200/60\">Basic Info</h1></div><!-- Form --><div class=\"max-w-sm mx-auto\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = createProfileForm().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"text-center mt-4\"><div class=\"text-sm text-slate-400\">Already have an account? <a class=\"font-medium text-purple-500 hover:text-purple-400 transition duration-150 ease-in-out\" href=\"/login\">Sign in</a></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = auth.RecoveryFooter().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div></section></main>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func createProfileForm() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form><div class=\"space-y-4\"><div><label class=\"block text-sm text-slate-300 font-medium mb-1\" for=\"email\">Email <span class=\"text-rose-500\">*</span></label> <input id=\"email\" class=\"form-input w-full\" type=\"username\" placeholder=\"mkiay\" required></div></div><div class=\"mt-6\"><button class=\"btn text-sm text-slate-900 bg-gradient-to-r from-white/80 via-white to-white/80 hover:bg-white w-full shadow-sm group\"><span class=\"mr-1\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"18\" height=\"17\" fill=\"none\" viewBox=\"0 0 24 24\"><path fill=\"#000\" d=\"M2.45759 5.40673C2.32495 5.61793 2.41154 5.88824 2.62359 6.0195L11.4735 11.498C11.796 11.6977 12.2037 11.6977 12.5262 11.498L21.3764 6.01941C21.5884 5.88814 21.675 5.61782 21.5423 5.40662C21.0117 4.56167 20.0714 4 19 4H5C3.92853 4 2.98825 4.56171 2.45759 5.40673Z\"></path><path fill=\"#000\" d=\"M22 8.88312C22 8.49137 21.5699 8.25179 21.2368 8.45799L13.579 13.1986C12.6114 13.7975 11.3883 13.7975 10.4208 13.1986L2.76318 8.45812C2.43008 8.25192 2 8.4915 2 8.88326V17C2 18.6569 3.34315 20 5 20H19C20.6569 20 22 18.6569 22 17V8.88312Z\"></path></svg></span> Send Confirmation Code <span class=\"tracking-normal text-slate-700 group-hover:translate-x-0.5 transition-transform duration-150 ease-in-out ml-1\">-&gt;</span></button></div></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func CreateCredential(rpName, rpId, challenge string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_CreateCredential_c37e`,
		Function: `function __templ_CreateCredential_c37e(rpName, rpId, challenge){const publicKeyCredentialCreationOptions = {
		challenge: Uint8Array.from(challenge, c => c.charCodeAt(0)),
		rp: {
			name: rpName,
			id: rpId,
		},
		user: {
			id: Uint8Array.from(
				"UZSL85T9AFC", c => c.charCodeAt(0)),
			name: "lee@webauthn.guide",
			displayName: "Lee",
		},
		pubKeyCredParams: [{alg: -7, type: "public-key"}],
		authenticatorSelection: {
			authenticatorAttachment: "cross-platform",
		},
		timeout: 60000,
		attestation: "direct"
	};

	const credential = navigator.credentials.create({
		publicKey: publicKeyCredentialCreationOptions
	}).then((credential) => {
		console.log(credential);
	});
}`,
		Call:       templ.SafeScript(`__templ_CreateCredential_c37e`, rpName, rpId, challenge),
		CallInline: templ.SafeScriptInline(`__templ_CreateCredential_c37e`, rpName, rpId, challenge),
	}
}
