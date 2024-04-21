import Cookies from 'js-cookie';
import React from 'react';
import { Navigate, Outlet } from 'react-router-dom';

const ProtectedRoute = (): React.ReactElement => {
  const bearerToken = Cookies.get('token')
  const auth = { token: bearerToken?.split(' ')[1] };

  return auth.token ? <Outlet /> : <Navigate to="/login" />;
};

export default ProtectedRoute;
