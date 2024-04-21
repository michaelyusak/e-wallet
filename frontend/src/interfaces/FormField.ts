export interface IFormField {
  type: 'email' | 'password' | 'text' | 'button' | 'currency';
  name: string;
  label?: string;
  placeholder?: string;
  isForSignIn?: boolean;
  isRequired?: boolean;
  balance?: string;
  onChange?: () => void;
}
