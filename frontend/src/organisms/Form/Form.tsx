import React, { ChangeEvent, FormEvent, useState } from 'react';

import { IFormField } from '../../interfaces/FormField';
import { Input } from '../../molecules/Input/Input';
import * as S from './Form.Styles';

type FormProps = {
  formFields: IFormField[];
  gap?: string;
  onSubmit: (inputValues: {
    [key: string]: { value: string; isError: boolean };
  }) => Promise<void>;
};

const Form = ({ formFields, gap, onSubmit }: FormProps): React.ReactElement => {
  const [inputValues, setInputVal] = useState<{
    [key: string]: { value: string; isError: boolean };
  }>({});

  // const [isAllOk, setAllOk] = useState<boolean>(false);

  // figure out how to undisable submit button

  async function handleOnSubmit(e: FormEvent<HTMLFormElement>) {
    e.preventDefault();

    await onSubmit(inputValues);
  }

  function handleInputChange(name: string, value: string, isError: boolean) {
    setInputVal((prevInputValues) => ({
      ...prevInputValues,
      [name]: { value, isError },
    }));
  }

  function handleInputValueChange(
    name: string,
    e: ChangeEvent<HTMLInputElement>,
    isError: boolean,
  ) {
    handleInputChange(name, e.target.value, isError);
  }

  return (
    <S.FormFields $Gap={gap} onSubmit={(e) => handleOnSubmit(e)}>
      {formFields.map((formField) => (
        <>
          {formField.type === 'button' ? (
            <S.Button type="submit" value={formField.placeholder}></S.Button>
          ) : (
            <Input
              value={inputValues[formField.name]?.value}
              handleOnChange={(name, e, isError) =>
                handleInputValueChange(name, e, isError)
              }
              formField={formField}
            ></Input>
          )}
        </>
      ))}
    </S.FormFields>
  );
};

export default Form;
