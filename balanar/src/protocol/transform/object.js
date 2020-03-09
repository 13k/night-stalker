import _ from "lodash";

const TRANSFORM_KEY = "$t";
const TRANSFORM_KEY_PLACEHOLDER = "$t";

export function propertyPath(path) {
  if (_.isString(path)) {
    path = _.split(path, ".");
  } else if (!_.isArray(path)) {
    throw new TypeError(`invalid path type ${path} (${typeof path})`);
  }

  if (path.length === 0) {
    throw new Error("empty path");
  }

  if (path[0] !== TRANSFORM_KEY_PLACEHOLDER) {
    path.unshift(TRANSFORM_KEY);
  }

  return _.map(path, seg => (seg === TRANSFORM_KEY_PLACEHOLDER ? TRANSFORM_KEY : seg));
}

export function property(path) {
  return _.property(propertyPath(path));
}

export function propertyMatches(path, value) {
  if (_.isPlainObject(path)) {
    return _.matches({ [TRANSFORM_KEY]: path });
  }

  return _.matchesProperty(propertyPath(path), value);
}

export function get(object, path, defaultValue) {
  return _.get(object, propertyPath(path), defaultValue);
}

export function set(object, path, value) {
  _.set(object, propertyPath(path), value);
  return object;
}

export function bindGet(path) {
  return _.chain(get)
    .partialRight(path)
    .unary()
    .value();
}

export function transformProperty(object, path, newPath, transformation) {
  if (_.isFunction(newPath)) {
    transformation = newPath;
    newPath = path;
  }

  return set(object, newPath, transformation(_.get(object, path)));
}
