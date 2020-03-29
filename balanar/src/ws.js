import Vue from "vue";
import _ from "lodash";

const EVENTS = ["open", "close", "error", "message"];

const STATE_CONNECTING = 0; // Socket has been created. The connection is not yet open.
const STATE_OPEN = 1; // The connection is open and ready to communicate.
const STATE_CLOSING = 2; // The connection is in the process of closing.
const STATE_CLOSED = 3; // The connection is closed or couldn't be opened.

const log = Vue.log({ context: { location: "ws" } });

class WSEvent extends Event {
  constructor(name, payload) {
    super(name);
    _.assign(this, payload || {});
  }
}

class WSError extends Error {
  constructor(message) {
    super(message);
    this.name = "WSError";
  }
}

export default class WS extends EventTarget {
  constructor(url) {
    super();

    this.url = url;
    this.socket = null;
  }

  get connecting() {
    return _.get(this, "socket.readyState") === STATE_CONNECTING;
  }

  get ready() {
    return _.get(this, "socket.readyState") === STATE_OPEN;
  }

  get closing() {
    return _.get(this, "socket.readyState") === STATE_CLOSING;
  }

  get closed() {
    return !this.socket || _.get(this, "socket.readyState") === STATE_CLOSED;
  }

  connect() {
    log.debug("connecting", this.url);
    this.socket = new WebSocket(this.url);

    _.each(EVENTS, event => {
      this.socket.addEventListener(event, this[`_on${event}`].bind(this));
    });
  }

  send(data) {
    if (!this.ready) {
      throw WSError("websocket not ready");
    }

    this.socket.send(data);
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
    log.debug("message", ev);
    const payload = { ev, data: ev.data, binaryType: this.socket.binaryType };
    this.dispatchEvent(new WSEvent("message", payload));
  }
}
