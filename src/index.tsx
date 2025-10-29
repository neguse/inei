import { Hono } from "hono";
import { renderer } from "./renderer";
import { Script } from "vite-ssr-components/hono";

const app = new Hono();

app.use(renderer);

app.get("/", (c) => {
  return c.render(
    <>
      <div id="screen-container">
        <iframe src="screen.html" allow="autoplay"></iframe>
      </div>
      <div id="editor-container">
        <div id="editor"></div>
      </div>
      <Script src="/src/editor.tsx" />
    </>
  );
});

export default app;
