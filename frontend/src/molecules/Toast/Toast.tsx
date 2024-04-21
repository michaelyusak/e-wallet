import React from 'react';
import * as S from './Toast.Styles';

type ToastProps = {
  isSuccess: boolean;
  message: string;
  marginTop?: string;
  marginLeft?: string;
};

const Toast = ({
  isSuccess,
  message,
  marginTop,
  marginLeft,
}: ToastProps): React.ReactElement => {
  return (
    <S.Toast
      $isSucces={isSuccess}
      $marginTop={marginTop}
      $marginLeft={marginLeft}
    >
      {message}
    </S.Toast>
  );
};

export default Toast;
