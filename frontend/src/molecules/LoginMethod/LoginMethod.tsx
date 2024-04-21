import React from 'react';

import * as S from './LoginMethod.Styles';
import iconFacebook from '../../assets/img/Facebook.png';
import iconApple from '../../assets/img/apple.png';
import iconGoogle from '../../assets/img/google.png';

const LoginMethod = (): React.ReactElement => {
  return (
    <S.Methods>
      <p>or continue with</p>

      <div>
        <button>
          <img src={iconFacebook} alt=""></img>
        </button>
        <button>
          <img src={iconApple} alt=""></img>
        </button>
        <button>
          <img src={iconGoogle} alt=""></img>
        </button>
      </div>
    </S.Methods>
  );
};

export default LoginMethod;
