import React, {
  ChangeEvent,
  FormEvent,
  useContext,
  useEffect,
  useState,
} from 'react';

import * as S from './ProfileModal.Styles';
import { IProfile } from '../../interfaces/Profile';
import editProfilePictureIcon from '../../assets/img/edit-profile-picture.png';
import { Input } from '../../molecules/Input/Input';
import { IFormField } from '../../interfaces/FormField';
import { ToastContext } from '../../contexts/ToastData';
import { useAppDispatch } from '../../redux/ReduxHooks';
import {
  fetchUser,
} from '../../features/UserData/userSlice';

type ProfileModalProps = {
  profile: IProfile;
  refetchProfile: () => void;
  closeProfile: () => void;
};

const ProfileModal = ({
  profile,
  refetchProfile,
  closeProfile,
}: ProfileModalProps): React.ReactElement => {
  const emailFormfield: IFormField = {
    type: 'email',
    name: 'email',
    isRequired: true,
  };

  const fullNameFormfield: IFormField = {
    type: 'text',
    name: 'fullName',
    placeholder: 'Enter full name',
    isRequired: true,
  };

  const [isEditProfile, setEditProfile] = useState<boolean>(false);
  const [inputValues, setInputValues] = useState<{
    [key: string]: { value: string; isError: boolean };
  }>({
    [emailFormfield.name]: { value: profile.email, isError: false },
    [fullNameFormfield.name]: { value: profile.username, isError: false },
  });
  const [isSubmitDisabled, setSubmitDisabled] = useState<boolean>(true);

  function handleSetInputValues(
    key: string,
    e: ChangeEvent<HTMLInputElement>,
    isError: boolean,
  ) {
    setInputValues((prevList) => ({
      ...prevList,
      [key]: { value: e.target.value, isError: isError },
    }));
  }

  const dispatch = useAppDispatch();

  function handleFormChange() {
    if (
      inputValues[emailFormfield.name].isError ||
      inputValues[fullNameFormfield.name].isError
    ) {
      setSubmitDisabled(true);
      return;
    }

    if (
      inputValues[emailFormfield.name].value === profile.email ||
      inputValues[fullNameFormfield.name].value === profile.username
    ) {
      setSubmitDisabled(true);
      return;
    }

    setSubmitDisabled(false);
  }

  function handleSetEditProfile(val: boolean) {
    setInputValues({
      [emailFormfield.name]: { value: profile.email, isError: false },
      [fullNameFormfield.name]: { value: profile.username, isError: false },
    });

    setEditProfile(val);
  }

  const { setToast } = useContext(ToastContext);

  async function handleUpdateUserData(e: FormEvent<HTMLFormElement>) {
    e.preventDefault();

    const email = inputValues[emailFormfield.name].value;
    const name = inputValues[fullNameFormfield.name].value;

    const url = 'http://localhost:8080/users';
    const options = {
      method: 'PATCH',
      headers: {
        'Content-Type': 'application/json',
        Authorization: localStorage.getItem('token') || '',
      },
      body: JSON.stringify({
        email: email,
        name: name,
      }),
    };

    try {
      const response = await fetch(url, options);
      const responseData = await response.json();

      if (!response.ok) {
        throw new Error(`Update user data failed: ${responseData.message}`);
      }

      handleShowToast(true, 'Profile Updated');

      refetchProfile();
      closeProfile();
    } catch (error) {
      console.log(error);

      handleShowToast(false, 'Failed Updating Profile');
    }
  }

  async function handleChangePicture(e: ChangeEvent<HTMLInputElement>) {
    e.preventDefault();

    const files = e.target.files;

    if (files?.length !== 1) {
      return;
    }

    const url = 'http://localhost:8080/users/pictures';
    const token = localStorage.getItem('token') || '';

    const formData = new FormData();
    formData.append('file', files[0]);

    const options = {
      method: 'PATCH',
      headers: { Authorization: token },
      body: formData,
    };

    try {
      const response = await fetch(url, options);
      const responseData = await response.json();

      if (!response.ok) {
        handleShowToast(false, 'Update Profile Picture Failed');
        throw new Error(responseData.message);
      }

      handleShowToast(true, 'Profile Picture Updated');
      dispatch(fetchUser())
    } catch (error) {
      console.log(error)
    }
  }

  useEffect(() => {
    handleFormChange();
    return () => {};
  }, [inputValues]);

  function handleShowToast(isSuccess: boolean, message: string) {
    setToast((prevToast) => ({
      ...prevToast,
      isSuccess: isSuccess,
      message: message,
      marginLeft: '39.5%',
      marginTop: '5%',
      isVisible: true,
    }));

    setTimeout(
      () =>
        setToast((prevToast) => ({
          ...prevToast,
          isSuccess: undefined,
          message: '',
          isVisible: false,
          marginLeft: undefined,
          marginTop: undefined,
        })),
      5000,
    );
  }

  return (
    <S.ProfileView>
      <S.ProfileUnderlay onClick={() => closeProfile()}></S.ProfileUnderlay>

      <S.ProfileCard $isEdit={isEditProfile}>
        <div>
          <img src={profile.picture} alt=""></img>

          <div>
            <label htmlFor="inputPicture">
              <img src={editProfilePictureIcon} alt=""></img>
            </label>

            <input
              type="file"
              name="file"
              id="inputPicture"
              onChange={(e) => handleChangePicture(e)}
            ></input>
          </div>
        </div>

        {!isEditProfile ? (
          <S.ProfileData>
            <div>
              <h1>{profile.username}</h1>

              <p>{profile.email}</p>
            </div>

            <button onClick={() => handleSetEditProfile(true)}>
              Edit Profile
            </button>
          </S.ProfileData>
        ) : (
          <S.EditProfileForm onSubmit={(e) => handleUpdateUserData(e)}>
            <div>
              <h3>Email</h3>

              <Input
                formField={emailFormfield}
                value={inputValues[emailFormfield.name].value}
                color="#A7A3FF"
                handleOnChange={(name, e, isError) =>
                  handleSetInputValues(name, e, isError)
                }
              ></Input>
            </div>

            <div>
              <h3>Full Name</h3>

              <Input
                formField={fullNameFormfield}
                value={inputValues[fullNameFormfield.name].value}
                color="#A7A3FF"
                handleOnChange={(name, e, isError) =>
                  handleSetInputValues(name, e, isError)
                }
              ></Input>
            </div>

            <div>
              <input
                type="submit"
                value="Save"
                disabled={isSubmitDisabled}
              ></input>

              <button type="button" onClick={() => handleSetEditProfile(false)}>
                Cancel
              </button>
            </div>
          </S.EditProfileForm>
        )}
      </S.ProfileCard>
    </S.ProfileView>
  );
};

export default ProfileModal;
