import React, { ChangeEvent, useState } from 'react';

import * as S from './Input.Styles';
import { IFormField } from '../../interfaces/FormField';
import showPassword from '../../assets/img/invisible-1.png';

type InputProps = {
  formField: IFormField;
  value: string;
  color?: string;
  handleOnChange: (
    name: string,
    e: ChangeEvent<HTMLInputElement>,
    isError: boolean,
  ) => void;
};

export const Input = ({
  formField,
  value,
  color,
  handleOnChange,
}: InputProps): React.ReactElement => {
  const formatter = new Intl.NumberFormat('en-US');

  const [isError, setError] = useState<boolean>(false);

  const [currencyError, setCurrencyError] = useState<
    'insufficientBalance' | 'invalidAmount' | undefined
  >(undefined);

  const [isPasswordShowed, setPasswordShowed] = useState<boolean>(false);

  function handleSetPasswordShowed() {
    setPasswordShowed(!isPasswordShowed);
  }

  function handleCurrencyOnChange(e: ChangeEvent<HTMLInputElement>) {
    handleInputCurrencyError(e);

    const editedVal = e.target.value.replaceAll(',', '');

    if (!value) {
      value = '';
    }

    if (isNaN(+editedVal)) {
      e.target.value = value;
    } else {
      e.target.value = editedVal; 
    }

    handleChange(e);
  }

  function handleChange(e: ChangeEvent<HTMLInputElement>) {
    const isError = e.target.validationMessage === '' ? false : true;

    setError(isError);

    handleOnChange(formField.name, e, isError);
  }

  function handleInputCurrencyError(e: ChangeEvent<HTMLInputElement>) {
    const inputNumber = +e.target.value.replaceAll(',', '');

    if (formField.label === 'Transfer') {
      if (!formField.balance) {
        return;
      }

      if (inputNumber > +formField.balance) {
        setCurrencyError('insufficientBalance');
        setError(true);
        return;
      }

      if (inputNumber < 1000 || inputNumber > 50000000) {
        setCurrencyError('invalidAmount');
        setError(true);
        return;
      }
    }

    if (formField.label === 'Top Up') {
      if (inputNumber < 50000 || inputNumber > 10000000) {
        setCurrencyError('invalidAmount');
        setError(true);
        return;
      }
    }

    setCurrencyError(undefined);
    setError(false);
  }

  return (
    <>
      {formField.type === 'password' && formField.isForSignIn ? (
        <>
          <S.InputSignInPassword $isError={isError}>
            <div>
              <input
                type={isPasswordShowed ? 'text' : formField.type}
                name={formField.name}
                value={value}
                onChange={(e) => handleChange(e)}
                onBlur={(ErrorEvent) =>
                  ErrorEvent.target.validationMessage
                    ? setError(true)
                    : setError(false)
                }
                placeholder={formField.placeholder}
                required={formField.isRequired}
              ></input>

              <button type="button" onClick={() => handleSetPasswordShowed()}>
                <img src={showPassword} alt=""></img>
              </button>
            </div>

            <p>Forgot password</p>
          </S.InputSignInPassword>
        </>
      ) : (
        <>
          {formField.type === 'password' ? (
            <S.InputPassword $isError={isError}>
              <input
                type={isPasswordShowed ? 'text' : formField.type}
                name={formField.name}
                value={value}
                onChange={(e) => handleChange(e)}
                onBlur={(ErrorEvent) =>
                  ErrorEvent.target.validationMessage
                    ? setError(true)
                    : setError(false)
                }
                placeholder={formField.placeholder}
                required={formField.isRequired}
              ></input>

              <button type="button" onClick={() => handleSetPasswordShowed()}>
                <img src={showPassword} alt=""></img>
              </button>
            </S.InputPassword>
          ) : formField.type === 'currency' ? (
            <S.InputCurrency $Error={currencyError}>
              <div>
                <div>
                  <p>IDR</p>
                </div>

                <input
                  type={formField.type}
                  name={formField.name}
                  value={value ? formatter.format(+value) : ''}
                  onChange={(e) => handleCurrencyOnChange(e)}
                  placeholder={formField.placeholder}
                  required={formField.isRequired}
                ></input>
              </div>

              {currencyError === 'invalidAmount' && (
                <S.ErrorStatement>
                  {formField.label} Value must be between{' '}
                  {(formField.label === 'Transfer' &&
                    'IDR 1,000 - IDR 50,000,000') ||
                    (formField.label === 'Top Up' &&
                      'IDR 50,000 - IDR 10,000,000')}
                </S.ErrorStatement>
              )}

              {currencyError === 'insufficientBalance' && (
                <S.ErrorStatement>Balance are not enough</S.ErrorStatement>
              )}

              {formField.balance !== undefined && (
                <S.Balance>
                  Remaining Balance: IDR{' '}
                  {formatter.format(formField.balance ? +formField.balance : 0)}
                </S.Balance>
              )}
            </S.InputCurrency>
          ) : (
            <S.InputField $isError={isError} $fontColor={color}>
              <input
                type={formField.type}
                name={formField.name}
                value={value}
                onChange={(e) => handleChange(e)}
                onBlur={(ErrorEvent) =>
                  ErrorEvent.target.validationMessage !== ''
                    ? setError(true)
                    : setError(false)
                }
                placeholder={
                  formField.type === 'email'
                    ? 'Enter email'
                    : formField.placeholder
                }
                required={formField.isRequired}
              ></input>
            </S.InputField>
          )}
        </>
      )}
    </>
  );
};
