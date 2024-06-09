import React from 'react';
import { Route, Routes } from 'react-router-dom';
import Registration from './components/Registrataion';
import Login from './components/Login';

function App() {
  return (
    <>
      <Routes>
        <Route path='/' element={<Registration />} />
        <Route path='/login' element={<Login />} />
      </Routes>
    </>
  );
}

export default App;
