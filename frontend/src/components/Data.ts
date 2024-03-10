// Generated by https://quicktype.io

export interface Data {
  Info: Info;
  Commit: Char;
  Char: Char;
  Han: Han;
  Pair: Pair;
  Keys: Keys;
  Dist: Dist;
}

export interface Char {
  Count: number;
  Word: number;
  WordFirst: number;
  Collision: number;
}

export interface Dist {
  CodeLen: number[];
  WordLen: number[];
  Collision: number[];
  Finger: number[];
  Key: { [key: string]: number };
}

export interface Han {
  NotHan: string;
  NotHans: number;
  NotHanCount: number;
  Lack: string;
  Lacks: number;
  LackCount: number;
}

export interface Info {
  TextName: string;
  TextLen: number;
  DictName: string;
  DictLen: number;
  Single: boolean;
}

export interface Keys {
  Count: number;
  CodeLen: number;
  LeftHand: number;
  RightHand: number;
}

export interface Pair {
  Count: number;
  Equivalent: number;
  SameFinger: number;
  DoubleHit: number;
  TribleHit: number;
  SingleSpan: number;
  MultiSpan: number;
  Staggered: number;
  Disturb: number;
  LeftToLeft: number;
  LeftToRight: number;
  RightToLeft: number;
  RightToRight: number;
  DiffFinger: number;
  SameHand: number;
  DiffHand: number;
}
