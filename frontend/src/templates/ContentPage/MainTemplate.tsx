import React, { useContext, useEffect, useState } from 'react';
import SideNav from '../../organisms/SideNav/SideNav';
import dashboardIcon from '../../assets/img/monitor-dashboard.png';
import transactionsIcon from '../../assets/img/cash-multiple.png';
import transferIcon from '../../assets/img/Transfer.png';
import topupIcon from '../../assets/img/Top-Up.png';
import logOutIcon from '../../assets/img/exit-to-app.png';
import { INavItem } from '../../interfaces/NavItem';
import { Outlet } from 'react-router';
import * as S from './MainTemplate.Styles';
import ContentHeader from '../../organisms/ContentHeader/ContentHeader';
import { useAppDispatch, useAppSelector } from '../../redux/ReduxHooks';
import { fetchUser } from '../../features/UserData/userSlice';
import { IProfile } from '../../interfaces/Profile';
import profileIcon from '../../assets/img/user.png';
import ProfileModal from '../../organisms/ProfileModal/ProfileModal';
import { ToastContext } from '../../contexts/ToastData';
import Toast from '../../molecules/Toast/Toast';

const MainTemplate = (): React.ReactElement => {
  function handleLogout() {
    localStorage.removeItem('token');
  }

  const sideNavItems: INavItem[] = [
    {
      itemIconSrc: dashboardIcon,
      itemName: 'Dashboard',
      isDefault: true,
      path: '/main/dashboard',
    },
    {
      itemIconSrc: transactionsIcon,
      itemName: 'Transactions',
      path: '/main/transactions',
    },
    {
      itemIconSrc: transferIcon,
      itemName: 'Transfer',
      path: '/main/transfer',
    },
    {
      itemIconSrc: topupIcon,
      itemName: 'Top Up',
      path: '/main/topup',
    },
    {
      itemIconSrc: logOutIcon,
      itemName: 'Logout',
      onClick: () => handleLogout(),
      path: '/login',
    },
  ];

  const userData = useAppSelector((state) => state.user);

  const dispatch = useAppDispatch();

  const profilePictureUrl = `http://localhost:8080/public/profile_pictures/${userData.user?.profile_picture}`;

  useEffect(() => {
    dispatch(fetchUser());
    return () => {};
  }, []);

  const profile: IProfile = {
    username: userData.user?.name,
    email: userData.user?.email,
    picture: profilePictureUrl === undefined ? profileIcon : profilePictureUrl,
  };

  const [isViewProfile, setViewProfile] = useState<boolean>(false);

  function handleViewProfileModal() {
    setViewProfile(!isViewProfile);
  }

  const { toastData } = useContext(ToastContext);

  return (
    <>
      {toastData.isVisible && (
        <Toast
          message={toastData.message}
          isSuccess={toastData.isSuccess ?? false}
          marginLeft={toastData.marginLeft}
          marginTop={toastData.marginTop}
        ></Toast>
      )}

      <S.FullPage>
        <SideNav navItems={sideNavItems}></SideNav>

        <S.HeaderContainer>
          <ContentHeader
            profilePictureUrl={profile.picture}
            onViewProfile={() => handleViewProfileModal()}
          ></ContentHeader>
          <S.SeparatorLine></S.SeparatorLine>
        </S.HeaderContainer>
        <Outlet></Outlet>

        {isViewProfile && (
          <ProfileModal
            profile={profile}
            refetchProfile={() => dispatch(fetchUser())}
            closeProfile={() => handleViewProfileModal()}
          ></ProfileModal>
        )}
      </S.FullPage>
    </>
  );
};

export default MainTemplate;
