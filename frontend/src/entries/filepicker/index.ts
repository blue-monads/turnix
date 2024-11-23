
import FilePicker from "./FilePicker.svelte";

const r = (window as any)["__turnix_registry__"];

if (r) {
    r.RegisterFactory(`buildFilePicker`, (target: HTMLElement, opts: Record<string, any>) => {
        new FilePicker({
            target: target,
            props: opts as any,
        });
    });
} else {
    console.error(`'__turnix_registry__' not found @ register buildFilePicker`);
}

