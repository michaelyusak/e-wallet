import React, { useState } from 'react';

import * as S from './ContentHeader.Styles';
import checkProfileIcon from '../../assets/img/check-profile.png';
import profileLogoutIcon from '../../assets/img/profile-logout.png';
import { useNavigate } from 'react-router-dom';

type contentHeaderProps = {
  profilePictureUrl: string;
  onViewProfile: () => void;
};

const ContentHeader = ({
  profilePictureUrl,
  onViewProfile,
}: contentHeaderProps): React.ReactElement => {
  const navigate = useNavigate();

  const [isProfileClicked, setProfileMenu] = useState<boolean>(false);

  function handleSetProfileMenu() {
    setProfileMenu(!isProfileClicked);
  }

  function handleShowProfileModal() {
    setProfileMenu(!isProfileClicked);

    onViewProfile();
  }

  const handleLogout = () => {
    localStorage.removeItem('token');

    navigate('/login');
  };

  const currentPath = window.location.pathname;
  const section = currentPath.split('/').pop();

  return (
    <S.ContentHeader>
      <h2>
        {section === 'dashboard'
          ? 'Dashboard'
          : section === 'transactions'
          ? 'Transactions'
          : ''}
      </h2>

      <S.ProfileNav>
        <button onClick={() => handleSetProfileMenu()}>
          <img src={profilePictureUrl} alt=""></img>
        </button>

        {isProfileClicked && (
          <ul>
            <li>
              <button onClick={() => handleShowProfileModal()}>
                <img src={checkProfileIcon} alt=""></img>

                <p>Profile</p>
              </button>
            </li>

            <li>
                <button onClick={() => handleLogout()}>
                  <img src={profileLogoutIcon} alt=""></img>

                  <p>Logout</p>
                </button>
            </li>
          </ul>
        )}
      </S.ProfileNav>
    </S.ContentHeader>
  );
};

export default ContentHeader;
