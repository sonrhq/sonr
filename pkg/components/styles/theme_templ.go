// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package styles

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func DownloadIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\"><g fill=\"currentColor\"><path d=\"M11.5 14.293l-.86-.86 -2.8-2.8c-.2-.2-.52-.2-.71 0l-.01 0c-.2.19-.2.51 0 .7h0l3.79 3.79v0c.58.58 1.53.58 2.12 0l0-.01 3.79-3.8s0-.01 0-.01c.19-.2.19-.52 0-.71 -.01 0-.01 0-.01 0m-5.36 3.64l4.29-4c.39-.4 1.02-.4 1.41 0l-.36.35m-5.36 3.64v-1.21 -8.09c0-.28.22-.5.5-.5 .27 0 .5.22.5.5V13v1.2l.85-.86 2.79-2.8m-4.65 3.64l4.64-3.65m.7 0c-.2-.2-.52-.2-.71 0 0 0-.01 0-.01 0m.7 0h-.71\"></path><path d=\"M11.85 13.93l-.86-.86 -2.8-2.8c-.4-.4-1.03-.4-1.42-.01l-.14.25 -.01 0 .49.09 -.36-.36c-.4.39-.4 1.02-.001 1.41l3.79 3.79c.78.78 2.04.78 2.82 0l0-.01 3.79-3.8c.39-.4.39-1.03 0-1.42l-.71.7s0 0-.01 0l-3.8 3.79 -.01 0c-.4.39-1.03.39-1.42-.01l-3.8-3.8s0 0-.01 0l.13-.26 0-.01 -.5-.1 .35.35c0-.01 0-.01 0-.01l2.79 2.79 .85.85 .7-.71Zm-.02.71l4.29-4c.2-.21.52-.21.71-.02l.35-.36 -.36-.36 -.36.35 .7.7 .35-.36 .35-.36 -.36-.36c-.59-.59-1.54-.59-2.13 0l-4.28 3.98 .68.73Zm.15-.37v-1.21 -8.09s0 0 0 0 0-.01 0 0v8.08 1.2 1.2l.85-.86 .85-.86 2.79-2.8 -.71-.71 -2.8 2.79 -.86.85 .85.35V13 4.91c0-.56-.45-1-1-1 -.56 0-1 .44-1 1v8.08 1.2h1Zm-.2.39l4.64-3.65 -.62-.79 -4.65 3.64 .61.78Zm5.39-4.4c-.4-.4-1.03-.4-1.42-.001l.7.7c-.01 0-.01 0-.01-.001l.7-.71Zm-.36-.15h-.71v1h.7v-1Z\"></path><path d=\"M4 14.5c.27 0 .5.22.5.5v2c0 .82.67 1.5 1.5 1.5h12c.82 0 1.5-.68 1.5-1.5v-2c0-.28.22-.5.5-.5 .27 0 .5.22.5.5v2c0 1.38-1.12 2.5-2.5 2.5H6c-1.39 0-2.5-1.12-2.5-2.5v-2c0-.28.22-.5.5-.5Z\"></path><path d=\"M4 15s0-.01 0 0v2c0 1.1.89 2 2 2h12c1.1 0 2-.9 2-2v-2c0-.01-.01 0 0 0 0 0 0-.01 0 0v2c0 1.1-.9 2-2 2H6c-1.11 0-2-.9-2-2v-2c0-.01-.01 0 0 0Zm0-1c-.56 0-1 .44-1 1v2c0 1.65 1.34 3 3 3h12c1.65 0 3-1.35 3-3v-2c0-.56-.45-1-1-1 -.56 0-1 .44-1 1v2c0 .55-.45 1-1 1H6c-.56 0-1-.45-1-1v-2c0-.56-.45-1-1-1Z\"></path></g></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
