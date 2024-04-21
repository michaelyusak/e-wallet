export interface ISelector {
  initialValue: string;
  menus: string[];
  title: string | undefined;
  setVal: (val: string) => void;
}
