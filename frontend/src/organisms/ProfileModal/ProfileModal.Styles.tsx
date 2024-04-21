import styled from 'styled-components';

export const ProfileUnderlay = styled.section`
  background-color: #454545;
  opacity: 90%;
  width: 100%;
  height: 100%;
`;

export const ProfileView = styled.div`
  position: absolute;
  top: 0;
  z-index: 5;
  width: 100%;
  height: 100%;
`;

export const ProfileCard = styled.div<{ $isEdit?: boolean }>`
  position: absolute;
  z-index: 1;
  top: ${(props) => (props.$isEdit ? '20%' : '26%')};
  left: 38%;
  width: 24%;
  height: fit-content;
  background-color: white;
  border-radius: 9px;
  padding: 53px 44px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  gap: 50px;
  box-shadow: 0px 4px 4px 0px #00000040;

  > div:nth-child(1) {
    width: 130px;
    height: 130px;
    position: relative;
    margin: 0 auto;

    img {
      width: 130px;
      height: 130px;
      object-fit: cover;
      border-radius: 100%;
    }

    > div {
      position: absolute;
      z-index: 99;
      bottom: 0;
      right: 0;
      height: 25px;
      border: none;
      outline: none;
      background-color: transparent;
      cursor: pointer;

      input {
        display: none;
      }

      label {
        cursor: pointer;

        img {
          width: 33.48px;
          height: 25px;
        }
      }
    }
  }
`;

export const ProfileData = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  gap: 50px;

  div {
    text-align: center;

    h1 {
      color: #282828;
      line-height: 42px;
      font-size: 28px;
      font-weight: 600;
    }

    p {
      color: #95999e;
      line-height: 18px;
      font-size: 12px;
      font-weight: 600;
    }
  }

  button {
    padding: 15px 56px;
    border-radius: 9px;
    border: none;
    outline: none;
    background-color: #4d47c3;
    color: white;
    line-height: 24px;
    font-size: 16px;
    font-weight: 500;
    height: 60px;
    cursor: pointer;
  }
`;

export const EditProfileForm = styled.form`
  display: flex;
  flex-direction: column;
  gap: 30px;

  > div:nth-child(3) {
    display: flex;
    justify-content: space-between;

    button,
    input {
      width: 47%;
      height: 58px;
      border-radius: 9px;
      cursor: pointer;
      line-height: 24px;
      font-size: 16px;
      font-weight: 500;
      outline: none;
    }

    button {
      background-color: white;
      border: #4d47c3 solid 1px;
      color: #4d47c3;
    }

    input {
      background-color: #4d47c3;
      border: none;
      color: white;
    }

    input:disabled {
      background-color: #cac8ff;
    }
  }

  div:nth-child(1),
  div:nth-child(2) {
    display: flex;
    flex-direction: column;
    gap: 10px;

    h3 {
      line-height: 18px;
      font-size: 12px;
      font-weight: 600;
    }
  }
`;
