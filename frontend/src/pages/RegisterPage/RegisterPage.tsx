import React, { useContext } from 'react';

import * as S from './RegisterPage.Styles';
import imageRegister from '../../assets/img/Image.png';
import Form from '../../organisms/Form/Form';
import { IFormField } from '../../interfaces/FormField';
import LoginMethod from '../../molecules/LoginMethod/LoginMethod';
import { useNavigate } from 'react-router';
import { ToastContext } from '../../contexts/ToastData';
import Toast from '../../molecules/Toast/Toast';
import Cookies from 'js-cookie';

const RegisterPage = (): React.ReactElement => {
  const navigate = useNavigate();
  const registerFormFields: IFormField[] = [
    {
      type: 'email',
      name: 'email',
      isRequired: true,
    },
    {
      type: 'text',
      name: 'fullName',
      placeholder: 'Enter full name',
      isRequired: true,
    },
    {
      type: 'password',
      name: 'password',
      placeholder: 'Enter password',
      isRequired: true,
    },
    {
      type: 'password',
      name: 'confirmPassword',
      placeholder: 'Confirm password',
      isRequired: true,
    },
    {
      type: 'button',
      name: 'submit',
      placeholder: 'Register',
      isRequired: true,
    },
  ];

  const handleRegisteredUser = () => {
    navigate('/login');
  };

  const { toastData, setToast } = useContext(ToastContext);

    async function handleRegister(inputValues: {
      [key: string]: { value: string; isError: boolean };
    }) {
    const email: string = inputValues['email']?.value;
    const name: string = inputValues['fullName']?.value;
    const password: string = inputValues['password']?.value;
    const confirmationPassword: string = inputValues['confirmPassword']?.value;

    if (password != confirmationPassword) {
      throw new Error('password must match confirmation password');
    }

    const url = 'http://localhost:8080/users';
    const options = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        email: email,
        name: name,
        password: password,
      }),
    };

    try {
      const response = await fetch(url, options);
      const responseData = await response.json();

      if (response.status !== 201) {
        throw new Error(`Failed to register user ${responseData.message}`);
      }

      handleShowToast(true, 'Register successfully');

      const token: string = responseData.data.token;

      Cookies.set('token', `bearer ${responseData.data.token}`, { expires: new Date(new Date().getTime() + 1000 * 60)});

      navigate('/main/dashboard');
    } catch (error) {
      console.error(error);

      handleShowToast(false, 'Register failed');

      throw error;
    }
  }

  function handleShowToast(isSuccess: boolean, message: string) {
    setToast((prevToast) => ({
      ...prevToast,
      isSuccess: isSuccess,
      message: message,
      isVisible: true,
      marginLeft: '40%',
    }));

    setTimeout(
      () =>
        setToast((prevToast) => ({
          ...prevToast,
          isSuccess: undefined,
          message: '',
          isVisible: false,
          marginLeft: undefined,
        })),
      5000,
    );
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
              <b>Join Us!</b> Sea Wallet
            </p>
            <p>
              Already have an account? You can
              <button onClick={() => handleRegisteredUser()}>
                Login here !
              </button>
            </p>
          </S.PageDescription>

          <S.RegisterImg src={imageRegister} alt=""></S.RegisterImg>

          <S.RegisterForm>
            <h2>Register</h2>

            <Form
              formFields={registerFormFields}
              onSubmit={(inputValues) => handleRegister(inputValues)}
            ></Form>

            <S.MobileDescription>
              <p>
                If you don’t have an account register You can{' '}
                <button onClick={() => handleRegisteredUser()}>
                  Login here !
                </button>
              </p>
            </S.MobileDescription>

            <LoginMethod></LoginMethod>
          </S.RegisterForm>
        </S.PageContent>
      </S.FullPage>
    </>
  );
};

export default RegisterPage;
