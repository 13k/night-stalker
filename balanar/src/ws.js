import Vue from "vue";
import { each } from "lodash/collection";
import { assign, get } from "lodash/object";

const EVENTS = ["open", "close", "error", "message"];

const STATE_CONNECTING = 0; // Socket has been created. The connection is not yet open.
const STATE_OPEN = 1; // The connection is open and ready to communicate.
const STATE_CLOSING = 2; // The connection is in the process of closing.
const STATE_CLOSED = 3; // The connection is closed or couldn't be opened.

const log = Vue.log({ context: { location: "ws" } });

class WSEvent extends Event {
  constructor(name, data) {
    super(name);
    assign(this, data || {});
  }
}

class WSError extends Error {}

export default class WS extends EventTarget {
  constructor(url) {
    super();

    this.url = url;
    this.socket = null;
  }

  get connecting() {
    return get(this, "socket.readyState") === STATE_CONNECTING;
  }

  get ready() {
    return get(this, "socket.readyState") === STATE_OPEN;
  }

  get closing() {
    return get(this, "socket.readyState") === STATE_CLOSING;
  }

  get closed() {
    return !this.socket || get(this, "socket.readyState") === STATE_CLOSED;
  }

  connect() {
    log.debug("connecting", this.url);
    this.socket = new WebSocket(this.url);

    each(EVENTS, event => {
      this.socket.addEventListener(event, this[`_on${event}`].bind(this));
    });
  }

  send(data) {
    if (!this.ready) {
      throw WSError("websocket not ready");
    }

    this.socket.send(JSON.stringify(data));
  }

  _onopen(ev) {
    log.debug("open", ev);
    this.dispatchEvent(new WSEvent("open", { ev }));
  }

  _onclose(ev) {
    log.debug("close", ev);
    this.dispatchEvent(new WSEvent("close", { ev }));
    this.socket = null;
  }

  _onerror(ev) {
    log.debug("error", ev);
    this.dispatchEvent(new WSEvent("error", { ev }));
  }

  _onmessage(ev) {
    log.debug("message");
    const body = JSON.parse(ev.data || "null");
    this.dispatchEvent(new WSEvent("message", { ev, body }));
  }
}
