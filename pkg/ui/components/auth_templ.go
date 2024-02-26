// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func LoginHeader() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col w-full justify-center items-center space-y-2 gap-1 mx-auto text-center max-w-lg\"><div class=\"w-28 h-28 rounded-full bg-gradient-to-t from-teal-500 via-cyan-600 to-blue-700 border-2 border-stone-400/90 mt-3\"></div><sl-divider style=\"--spacing: 0.6rem;\"></sl-divider><h1 class=\"text-3xl font-semibold\">Login</h1><p class=\"text-stone-300 text-md text-center py-2 text-relaxed\">Begin your <sl-tooltip content=\"Access every IBC enabled blockchain with a secure Interchain Account.\" placement=\"top\"><span class=\"font-semibold text-stone-200\">Crypto Journey</span></sl-tooltip> seamlessly and access decentralized services <sl-tooltip content=\"End to End encryption from the Validator to the client. Multi-Party Computed Wallets. Secure communication channel over Matrix.\" placement=\"top\"><span class=\"font-semibold text-stone-200\">the Safe Way</span></sl-tooltip> by creating a <sl-tooltip content=\"Decentralized Identity representation on the Sonr Blockchain. Anonymous and Encrypted.\" placement=\"top\"><span class=\"font-semibold text-stone-200\">Sonr Identity</span></sl-tooltip>.</p></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func LoginFooter() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div slot=\"footer\" class=\"justify-between items-center w-full max-w-[12vw] mx-auto opacity-75 py-1.5 flex\"><sl-tooltip content=\"Need Help?\" placement=\"top\"><sl-icon-button name=\"life-preserver\" label=\"Settings\"></sl-icon-button></sl-tooltip> <sl-tooltip content=\"Documentation\" placement=\"top\"><sl-icon-button name=\"book-half\" label=\"Settings\"></sl-icon-button></sl-tooltip> <sl-tooltip content=\"Start Recovery\" placement=\"top\"><sl-icon-button name=\"tools\" label=\"Settings\"></sl-icon-button></sl-tooltip></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func RegisterHeader() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col w-full justify-center items-center space-y-2 gap-1 mx-auto text-center max-w-lg\"><div class=\"w-28 h-28 rounded-full bg-gradient-to-t from-teal-500 via-cyan-600 to-blue-700 border-2 border-stone-400/90 mt-3\"></div><sl-divider style=\"--spacing: 0.6rem;\"></sl-divider><h1 class=\"text-3xl font-semibold\">Register</h1><p class=\"text-stone-300 text-md text-center py-2 text-relaxed\">Begin your <sl-tooltip content=\"Access every IBC enabled blockchain with a secure Interchain Account.\" placement=\"top\"><span class=\"font-semibold text-stone-200\">Crypto Journey</span></sl-tooltip> seamlessly and access decentralized services <sl-tooltip content=\"End to End encryption from the Validator to the client. Multi-Party Computed Wallets. Secure communication channel over Matrix.\" placement=\"top\"><span class=\"font-semibold text-stone-200\">the Safe Way</span></sl-tooltip> by creating a <sl-tooltip content=\"Decentralized Identity representation on the Sonr Blockchain. Anonymous and Encrypted.\" placement=\"top\"><span class=\"font-semibold text-stone-200\">Sonr Identity</span></sl-tooltip>.</p></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func RegisterFooter() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div slot=\"footer\" class=\"justify-evenly items-center w-full max-w-[12vw] mx-auto opacity-75 py-1.5 flex\"><sl-tooltip content=\"Need Help?\" placement=\"top\"><sl-icon-button name=\"life-preserver\" label=\"Helpdesk\"></sl-icon-button></sl-tooltip> <sl-tooltip content=\"Documentation\" placement=\"top\"><sl-icon-button name=\"book-half\" label=\"Settings\"></sl-icon-button></sl-tooltip></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}