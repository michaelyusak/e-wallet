import { Dispatch, SetStateAction, createContext } from 'react';

export interface IToastData {
  isSuccess: boolean | undefined;
  message: string;
  isVisible?: boolean;
  marginLeft?: string;
  marginTop?: string;
}

interface IToastContext {
  toastData: IToastData;
  setToast: Dispatch<SetStateAction<IToastData>>;
}

const initialToastData: IToastData = {
  isSuccess: undefined,
  message: '',
};

export const ToastContext = createContext<IToastContext>({
  toastData: initialToastData,
  setToast: () => {},
});
