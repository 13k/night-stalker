import _ from "lodash";

/**
 * Normalizes the given data.
 * @param {ArrayBuffer | Blob | Body | Promise<ArrayBuffer |Blob | Body>} data Data to normalize.
 * @returns {Promise<Uint8Array>} Promise resolving to an Uint8Array.
 * @throws {TypeError} If data is not supported.
 */
async function normalizeData(data) {
  if (data instanceof Promise) {
    data = await data;
  }

  if (_.isArrayBuffer(data)) {
    return Promise.resolve(new Uint8Array(data));
  }

  if (_.isFunction(data.arrayBuffer)) {
    return data.arrayBuffer().then(arr => new Uint8Array(arr));
  }

  throw new TypeError("invalid data");
}

/**
 * Decodes the given data with the given decoder.
 * @function decode
 * @param {protobuf.Message} decoder Protobuf message class.
 * @param {ArrayBuffer|Blob|Body|Promise<ArrayBuffer|Blob|Body>} data Message data.
 * @returns {Promise<protobuf.Message>} Decoded message.
 * @throws {TypeError} if data is not supported.
 */
export const decode = (decoder, data) => normalizeData(data).then(decoder.decode);
