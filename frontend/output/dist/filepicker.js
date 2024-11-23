(function() {
  "use strict";
  const PUBLIC_VERSION = "5";
  if (typeof window !== "undefined")
    (window.__svelte || (window.__svelte = { v: /* @__PURE__ */ new Set() })).v.add(PUBLIC_VERSION);
  const TEMPLATE_USE_IMPORT_NODE = 1 << 1;
  let active_effect = null;
  var first_child_getter;
  // @__NO_SIDE_EFFECTS__
  function get_first_child(node) {
    return first_child_getter.call(node);
  }
  function create_fragment_from_html(html) {
    var elem = document.createElement("template");
    elem.innerHTML = html;
    return elem.content;
  }
  function assign_nodes(start, end) {
    var effect = (
      /** @type {Effect} */
      active_effect
    );
    if (effect.nodes_start === null) {
      effect.nodes_start = start;
      effect.nodes_end = end;
    }
  }
  // @__NO_SIDE_EFFECTS__
  function template(content, flags) {
    var use_import_node = (flags & TEMPLATE_USE_IMPORT_NODE) !== 0;
    var node;
    var has_start = !content.startsWith("<!>");
    return () => {
      if (node === void 0) {
        node = create_fragment_from_html(has_start ? content : "<!>" + content);
        node = /** @type {Node} */
        /* @__PURE__ */ get_first_child(node);
      }
      var clone = (
        /** @type {TemplateNode} */
        use_import_node ? document.importNode(node, true) : node.cloneNode(true)
      );
      {
        assign_nodes(clone, clone);
      }
      return clone;
    };
  }
  function append(anchor, dom) {
    if (anchor === null) {
      return;
    }
    anchor.before(
      /** @type {Node} */
      dom
    );
  }
  var root = /* @__PURE__ */ template(`<div>FilePicker</div>`);
  function FilePicker($$anchor) {
    var div = root();
    append($$anchor, div);
  }
  const r = window["__turnix_registry__"];
  if (r) {
    r.RegisterFactory(`buildFilePicker`, (target, opts) => {
      new FilePicker({
        target,
        props: opts
      });
    });
  } else {
    console.error(`'__turnix_registry__' not found @ register buildFilePicker`);
  }
})();
