import { basicSetup } from "codemirror";
import { EditorView, keymap } from "@codemirror/view";
import { go } from "@codemirror/lang-go";
import { oneDark } from "@codemirror/theme-one-dark";

fetch("/shader.kage")
  .then((res) => res.text())
  .then((code) => {
    const view = new EditorView({
      doc: code,
      extensions: [
        basicSetup,
        go(),
        oneDark,
        EditorView.theme({
          "&": { height: "100%" },
        }),
        keymap.of([
          {
            key: "Alt-Enter",
            run: (view) => {
              const code = view.state.doc.toString();
              const iframe = document.querySelector(
                'iframe[src="screen.html"]'
              ) as HTMLIFrameElement;
              if (iframe?.contentWindow) {
                (iframe.contentWindow as any).applyShader?.(code);
              }
              return true;
            },
          },
        ]),
      ],
      parent: document.getElementById("editor")!,
    });
  });
