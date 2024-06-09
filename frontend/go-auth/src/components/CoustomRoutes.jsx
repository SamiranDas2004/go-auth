import React from 'react'
import { Route,Router } from 'react-router-dom'
import Registration from './Registrataion'
import Login from './Login'
function CoustomRoutes() {
  return (
    <div>
          <Router>
      <Route path='/' element={ <Registration/>}/>
        <Route path='/login' element={<Login/>}/>
      </Router>
    </div>
  )
}

export default CoustomRoutes