import React, { useContext } from 'react';
import Cookies from 'js-cookie';

import imageSignIn from '../../assets/img/Saly-14.png';
import Form from '../../organisms/Form/Form';
import { IFormField } from '../../interfaces/FormField';
import LoginMethod from '../../molecules/LoginMethod/LoginMethod';

import * as S from './LoginPage.Styles';
import { useNavigate } from 'react-router';
import { ToastContext } from '../../contexts/ToastData';
import Toast from '../../molecules/Toast/Toast';

const LoginPage = (): React.ReactElement => {
  const navigate = useNavigate();
  const signInFormFields: IFormField[] = [
    {
      type: 'email',
      name: 'email',
      isRequired: true,
    },
    {
      type: 'password',
      name: 'password',
      placeholder: 'Password',
      isForSignIn: true,
      isRequired: true,
    },
    {
      type: 'button',
      name: 'submit',
      placeholder: 'Login',
    },
  ];
  const { toastData, setToast } = useContext(ToastContext);

  async function handleLogin(inputValues: {
    [key: string]: { value: string; isError: boolean };
  }) {
    const email: string = inputValues['email']?.value;
    const password: string = inputValues['password']?.value;

    const url = 'http://localhost:8080/users/login';
    const options = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        email: email,
        password: password,
      }),
    };

    try {
      const response = await fetch(url, options);
      const responseData = await response.json();

      if (!response.ok) {
        throw new Error(`Failed to sign in ${responseData.message}`);
      }

      Cookies.set('token', `bearer ${responseData.data.token}`, { expires: new Date(new Date().getTime() + 1000 * 60)});

      handleShowToast(true, 'Login successfully');

      navigate('/main/dashboard');
    } catch (error) {
      console.log(error);

      handleShowToast(false, 'Login failed');

      throw error;
    }
  }

  function handleShowToast(isSuccess: boolean, message: string) {
    setToast({
      isSuccess: isSuccess,
      message: message,
      isVisible: true,
    });

    setTimeout(
      () =>
        setToast((prevToast) => ({
          ...prevToast,
          isSuccess: undefined,
          message: '',
          isVisible: false,
        })),
      5000,
    );
  }

  function handleNewUser() {
    navigate('/register');
  }

  return (
    <>
      {toastData.isVisible && (
        <Toast
          message={toastData.message}
          isSuccess={toastData.isSuccess ?? false}
          marginLeft="40%"
        ></Toast>
      )}

      <S.FullPage>
        <S.Header>
          <S.Logo>Sea Wallet</S.Logo>
        </S.Header>

        <S.PageContent>
          <S.PageDescription>
            <p>
              <b>Sign in to</b> Sea Wallet
            </p>
            <p>
              If you don’t have an account register You can
              <button onClick={() => handleNewUser()}>Register here !</button>
            </p>
          </S.PageDescription>

          <S.LoginImg src={imageSignIn} alt=""></S.LoginImg>

          <S.SignInForm>
            <h2>Sign in</h2>

            <Form
              formFields={signInFormFields}
              onSubmit={(inputValues) => handleLogin(inputValues)}
            ></Form>

            <S.MobileDescription>
              <p>
                If you don’t have an account register You can{' '}
                <button onClick={() => handleNewUser()}>Register here !</button>
              </p>
            </S.MobileDescription>

            <LoginMethod></LoginMethod>
          </S.SignInForm>
        </S.PageContent>
      </S.FullPage>
    </>
  );
};

export default LoginPage;
