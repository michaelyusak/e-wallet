export interface INavItem {
  itemIconSrc: string;
  itemName: string;
  isDefault?: boolean;
  path: string;
  onClick?: () => void;
}
