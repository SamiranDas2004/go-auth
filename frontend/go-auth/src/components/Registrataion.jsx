import React, { useState } from 'react'
import axios from 'axios'
function Registrataion() {

    const [user,setUser]=useState({
        username:"",
        email:"",
        password:""
    })

    const handelSignUp=async()=>{
        try {
            const response=await axios.post("http://localhost:4000/registers",user)
            if (!response) {
                console.log("error in post");
            }
            console.log(response.data);
        } catch (error) {
            throw new Error(error.messsage)
        }
    }
console.log(user);


  return (
    <>
    <form className="flex flex-col items-center justify-center h-screen">
      <div className="md:flex md:items-center mb-6">
        <div className="md:w-1/3">
          <label className="block text-gray-500 font-bold md:text-right mb-1 md:mb-0 pr-4" htmlFor="inline-full-name">
            Full Name
          </label>
        </div>
        <div className="md:w-2/3">
          <input 
            className="bg-gray-200 appearance-none border-2 border-gray-200 rounded w-full py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-purple-500" 
            id="inline-full-name" 
            type="text" 
            placeholder="User Name" 
            value={user.username} 
            onChange={(e) => setUser({ ...user, username: e.target.value })}
          />
        </div>
      </div>
      <div className="md:flex md:items-center mb-6">
        <div className="md:w-1/3">
          <label className="block text-gray-500 font-bold md:text-right mb-1 md:mb-0 pr-4" htmlFor="inline-email">
            Email
          </label>
        </div>
        <div className="md:w-2/3">
          <input 
            className="bg-gray-200 appearance-none border-2 border-gray-200 rounded w-full py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-purple-500" 
            id="inline-email" 
            type="text" 
            placeholder="Email" 
            value={user.email} 
            onChange={(e) => setUser({ ...user, email: e.target.value })}
          />
        </div>
      </div>
      <div className="md:flex md:items-center mb-6">
        <div className="md:w-1/3">
          <label className="block text-gray-500 font-bold md:text-right mb-1 md:mb-0 pr-4" htmlFor="inline-password">
            Password
          </label>
        </div>
        <div className="md:w-2/3">
          <input 
            className="bg-gray-200 appearance-none border-2 border-gray-200 rounded w-full py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-purple-500" 
            id="inline-password" 
            type="password" 
            placeholder="******************" 
            value={user.password} 
            onChange={(e) => setUser({ ...user, password: e.target.value })}
          />
        </div>
      </div>
      <button 
        
        className=' font-3xl bold'
        onClick={handelSignUp}>signup</button>
      </form>
      <div>
       
      </div>
      </>
      
  )
}

export default Registrataion