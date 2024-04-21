import React from 'react';
import { Navigate, createBrowserRouter } from 'react-router-dom';
import LoginPage from '../pages/LoginPage/LoginPage';
import RegisterPage from '../pages/RegisterPage/RegisterPage';
import DashboardPage from '../pages/DashboardPage/DashboardPage';
import TransactionsPage from '../pages/TransactionPage/TransactionsPage';
import MainTemplate from '../templates/ContentPage/MainTemplate';
import ProtectedRoute from './ProtectedRoute';

const router = createBrowserRouter([
  {
    path: '*',
    element: <Navigate to="/main"></Navigate>,
  },
  {
    path: '/login',
    element: <LoginPage></LoginPage>,
  },
  {
    path: '/register',
    element: <RegisterPage></RegisterPage>,
  },
  {
    path: '/main',
    element: <ProtectedRoute></ProtectedRoute>,
    children: [
      {
        path: '/main',
        element: <MainTemplate></MainTemplate>,
        children: [
          {
            path: '/main/',
            element: <Navigate to="/main/dashboard"></Navigate>,
          },
          {
            path: '/main/dashboard',
            element: <DashboardPage></DashboardPage>,
          },
          {
            path: '/main/transactions',
            element: <TransactionsPage></TransactionsPage>,
          },
        ],
      },
    ],
  },
]);

export default router;
