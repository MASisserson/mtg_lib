interface ImageURIs {
  small: string;
  normal: string;
  large: string;
  png: string;
  art_crop: string;
  border_crop: string;
}

interface ColorSet {
  colors: Array<string>;
  white: boolean;
  blue: boolean;
  black: boolean;
  red: boolean;
  green: boolean;
}

export interface Card {
  id: number;
  name: string;
  wiz_id: string;
  set: string;
  image_uris: ImageURIs;
  oracle_text: string;
  type_line: string;
  color_set: ColorSet;
  mana_value: number;
  mana_cost: string;
  power: string;
  toughness: string;
  loyalty: string;
  quantity: number;
}