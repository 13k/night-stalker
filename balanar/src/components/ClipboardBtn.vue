<template>
  <v-btn
    ref="btn"
    :title="title"
    icon
    small
  >
    <slot>
      <v-icon small>
        mdi-content-copy
      </v-icon>
    </slot>
  </v-btn>
</template>

<script>
import Clipboard from "clipboard";

export default {
  name: "ClipboardBtn",

  props: {
    content: {
      type: [HTMLElement, String, Number, Function],
      required: true,
    },
    action: {
      type: String,
      default: "copy",
    },
    container: {
      type: HTMLElement,
      default: () => document.body,
    },
    title: {
      type: String,
      default: "Copy to clipboard",
    },
    success: {
      type: Function,
      default: () => {},
    },
    error: {
      type: Function,
      default: () => {},
    },
  },

  computed: {
    text() {
      let value;

      switch (typeof this.content) {
        case "string":
        case "number":
        case "bigint":
          value = this.content.toString();
          break;
        case "function":
          value = this.content();
          break;
        case "object":
          if (this.content instanceof HTMLInputElement) {
            value = this.content.value;
          } else if (this.content instanceof HTMLTextAreaElement) {
            value = this.content.value;
          } else if (this.content instanceof HTMLElement) {
            value = this.content.innerText;
          } else {
            throw "Invalid object type for 'content' property";
          }
          break;
        default:
          throw `Invalid type ${typeof this.content} for 'content' property`;
      }

      return value;
    },
  },

  watch: {
    content(val) {
      this.$log.debug("content()", val);
      this.destroyClipboard();
      this.createClipboard();
    },
  },

  mounted() {
    this.createClipboard();
  },

  destroyed() {
    this.destroyClipboard();
  },

  methods: {
    destroyClipboard() {
      if (this.clipboard != null) {
        this.clipboard.destroy();
      }
    },
    createClipboard() {
      const $vm = this;

      this.clipboard = new Clipboard(this.$refs.btn.$el, {
        text: () => $vm.text,
        action: () => $vm.action,
        container: this.container,
      });

      if (this.success) {
        this.clipboard.on("success", this.success);
      }

      if (this.error) {
        this.clipboard.on("error", this.error);
      }
    },
  },
};
</script>
