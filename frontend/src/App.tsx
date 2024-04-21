import React, { useState } from 'react';
import { RouterProvider } from 'react-router-dom';
import router from './router';
import { IToastData, ToastContext } from './contexts/ToastData';
import { Provider } from 'react-redux';
import store from './redux/Store';

const App: React.FC = () => {
  const [toastData, setToast] = useState<IToastData>({
    isSuccess: undefined,
    message: '',
  });

  return (
    <>
      <Provider store={store}>
        <ToastContext.Provider value={{ toastData, setToast }}>
          <RouterProvider router={router}></RouterProvider>
        </ToastContext.Provider>
      </Provider>
    </>
  );
};

export default App;
