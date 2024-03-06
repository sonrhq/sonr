// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package about

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func heroSection() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Hero --><section class=\"relative\"><!-- Radial gradient --><div class=\"absolute flex items-center justify-center top-0 -translate-y-1/2 left-1/2 -translate-x-1/2 pointer-events-none -z-10 w-[800px] aspect-square\" aria-hidden=\"true\"><div class=\"absolute inset-0 translate-z-0 bg-purple-500 rounded-full blur-[120px] opacity-30\"></div><div class=\"absolute w-64 h-64 translate-z-0 bg-purple-400 rounded-full blur-[80px] opacity-70\"></div></div><!-- Particles animation --><div class=\"absolute inset-0 h-96 -z-10\" aria-hidden=\"true\"><canvas data-particle-animation data-particle-quantity=\"15\"></canvas></div><!-- Illustration --><div class=\"md:block absolute left-1/2 -translate-x-1/2 -mt-16 blur-2xl opacity-90 pointer-events-none -z-10\" aria-hidden=\"true\"><img src=\"https://cdn.sonr.build/images/page-illustration.svg\" class=\"max-w-none\" width=\"1440\" height=\"427\" alt=\"Page Illustration\"></div><div class=\"max-w-6xl mx-auto px-4 sm:px-6\"><div class=\"pt-32 pb-10 md:pt-40\"><!-- Hero content --><div class=\"text-center\"><div class=\"inline-flex font-medium bg-clip-text text-transparent bg-gradient-to-r from-purple-500 to-purple-200 pb-3\">The folks who are building</div><h1 class=\"h1 bg-clip-text text-transparent bg-gradient-to-r from-slate-200/60 via-slate-200 to-slate-200/60 pb-6\">The Internet's Driver's License</h1><!-- Rings illustration --><div class=\"inline-flex items-center justify-center relative\"><!-- Particles animation --><div class=\"absolute inset-0 -z-10\" aria-hidden=\"true\"><canvas data-particle-animation data-particle-quantity=\"10\"></canvas></div><div class=\"inline-flex [mask-image:_radial-gradient(circle_at_bottom,transparent_15%,black_70%)]\"><img src=\"https://cdn.sonr.build/images/about-illustration.svg\" width=\"446\" height=\"446\" alt=\"About illustration\"></div><img class=\"absolute mt-[30%] drop-shadow-lg animate-float\" src=\"https://cdn.sonr.build/images/hyperauth.svg\" width=\"72\" height=\"72\" alt=\"About icon\"></div></div></div></div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}