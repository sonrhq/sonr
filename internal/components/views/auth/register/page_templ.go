// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package register

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/labstack/echo/v4"
	"github.com/sonrhq/sonr/internal/components/base"
)

func Page(c echo.Context) templ.Component {
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
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<main class=\"grow\"><section class=\"relative\"><!-- Illustration --><div class=\"md:block absolute left-1/2 -translate-x-1/2 -mt-36 blur-2xl opacity-70 pointer-events-none -z-10\" aria-hidden=\"true\"><img src=\"https://cdn.sonr.build/images/auth-illustration.svg\" class=\"max-w-none\" width=\"1440\" height=\"450\" alt=\"Page Illustration\"></div><div class=\"relative max-w-6xl mx-auto px-4 sm:px-6\"><div class=\"pt-32 pb-12 md:pt-40 md:pb-20\"><!-- Page header --><div class=\"max-w-3xl mx-auto text-center pb-12\"><!-- Logo --><div class=\"mb-5\"><a class=\"inline-flex\" href=\"/\"><div class=\"relative flex items-center justify-center w-16 h-16 border border-transparent rounded-2xl shadow-2xl [background:linear-gradient(theme(colors.slate.900),_theme(colors.slate.900))_padding-box,_conic-gradient(theme(colors.slate.400),_theme(colors.slate.700)_25%,_theme(colors.slate.700)_75%,_theme(colors.slate.400)_100%)_border-box] before:absolute before:inset-0 before:bg-slate-800/30 before:rounded-2xl\"><img class=\"relative\" src=\"https://cdn.sonr.build/images/logo.svg\" width=\"42\" height=\"42\" alt=\"Stellar\"></div></a></div><!-- Page title --><h1 class=\"h2 bg-clip-text text-transparent bg-gradient-to-r from-slate-200/60 via-slate-200 to-slate-200/60\">Register Account</h1></div><!-- Form --><div class=\"max-w-sm mx-auto\"><form><div class=\"space-y-4\"><div><label class=\"block text-sm text-slate-300 font-medium mb-1\" for=\"full-name\">Username <span class=\"text-rose-500\">*</span></label> <input id=\"full-name\" class=\"form-input w-full\" type=\"text\" placeholder=\"mrossi\" required></div><div><label class=\"block text-sm text-slate-300 font-medium mb-1\" for=\"email\">Email <span class=\"text-rose-500\">*</span></label> <input id=\"email\" class=\"form-input w-full\" type=\"email\" placeholder=\"markrossi@company.com\" required></div><div><label class=\"block text-sm text-slate-300 font-medium mb-1\" for=\"referrer\">Where did you hear about us? <span class=\"text-rose-500\">*</span></label> <select id=\"referrer\" class=\"form-select text-sm py-2 w-full\" required><option>Google</option> <option>Medium</option> <option>GitHub</option></select></div></div><div class=\"mt-6\"><button class=\"btn text-sm text-slate-900 bg-gradient-to-r from-white/80 via-white to-white/80 hover:bg-white w-full shadow-sm group\">Sign Up <span class=\"tracking-normal text-slate-700 group-hover:translate-x-0.5 transition-transform duration-150 ease-in-out ml-1\">-&gt;</span></button></div></form><div class=\"text-center mt-4\"><div class=\"text-sm text-slate-400\">Already have an account? <a class=\"font-medium text-purple-500 hover:text-purple-400 transition duration-150 ease-in-out\" href=\"signin.html\">Sign in</a></div></div><!-- Divider --><div class=\"flex items-center my-6\"><div class=\"border-t border-slate-800 grow mr-3\" aria-hidden=\"true\"></div><div class=\"text-sm text-slate-500 italic\">or</div><div class=\"border-t border-slate-800 grow ml-3\" aria-hidden=\"true\"></div></div><!-- Social login --><div class=\"flex space-x-3\"><button class=\"btn text-slate-300 hover:text-white transition duration-150 ease-in-out w-full group [background:linear-gradient(theme(colors.slate.900),_theme(colors.slate.900))_padding-box,_conic-gradient(theme(colors.slate.400),_theme(colors.slate.700)_25%,_theme(colors.slate.700)_75%,_theme(colors.slate.400)_100%)_border-box] relative before:absolute before:inset-0 before:bg-slate-800/30 before:rounded-full before:pointer-events-none h-9\"><span class=\"relative\"><span class=\"sr-only\">Continue with Twitter</span> <svg class=\"fill-current\" xmlns=\"http://www.w3.org/2000/svg\" width=\"14\" height=\"12\"><path d=\"m4.34 0 2.995 3.836L10.801 0h2.103L8.311 5.084 13.714 12H9.482L6.169 7.806 2.375 12H.271l4.915-5.436L0 0h4.34Zm-.635 1.155H2.457l7.607 9.627h1.165L3.705 1.155Z\"></path></svg></span></button> <button class=\"btn text-slate-300 hover:text-white transition duration-150 ease-in-out w-full group [background:linear-gradient(theme(colors.slate.900),_theme(colors.slate.900))_padding-box,_conic-gradient(theme(colors.slate.400),_theme(colors.slate.700)_25%,_theme(colors.slate.700)_75%,_theme(colors.slate.400)_100%)_border-box] relative before:absolute before:inset-0 before:bg-slate-800/30 before:rounded-full before:pointer-events-none h-9\"><span class=\"relative\"><span class=\"sr-only\">Continue with GitHub</span> <svg class=\"fill-current\" xmlns=\"http://www.w3.org/2000/svg\" width=\"16\" height=\"15\"><path d=\"M7.488 0C3.37 0 0 3.37 0 7.488c0 3.276 2.153 6.084 5.148 7.113.374.094.468-.187.468-.374v-1.31c-2.06.467-2.527-.936-2.527-.936-.375-.843-.843-1.124-.843-1.124-.655-.468.094-.468.094-.468.749.094 1.123.75 1.123.75.655 1.216 1.778.842 2.153.654.093-.468.28-.842.468-1.03-1.685-.186-3.37-.842-3.37-3.743 0-.843.281-1.498.75-1.966-.094-.187-.375-.936.093-1.965 0 0 .655-.187 2.059.749a6.035 6.035 0 0 1 1.872-.281c.655 0 1.31.093 1.872.28 1.404-.935 2.059-.748 2.059-.748.374 1.03.187 1.778.094 1.965.468.562.748 1.217.748 1.966 0 2.901-1.778 3.463-3.463 3.65.281.375.562.843.562 1.498v2.059c0 .187.093.468.561.374 2.996-1.03 5.148-3.837 5.148-7.113C14.976 3.37 11.606 0 7.488 0Z\"></path></svg></span></button></div></div></div></div></section></main>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = base.PageLayout().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
