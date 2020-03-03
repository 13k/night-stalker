export interface KeyValues {
  [propertyName: string]: string | KeyValues;
}

export function dump(obj: { [key: string]: any }, pretty: boolean): string;
export function stringify(obj: { [key: string]: any }, pretty: boolean): string;
export function parse(text: string): KeyValues;
