import React, { useState } from 'react'
import { Link } from 'react-router-dom'

import './login.css'
function Login() {
 
  interface loginuser {
    email : string ;
    password : string ;
  }
  const [email , setemail] = useState('');
  const [ password , setpassword] = useState('');

  const login = async (e : React.MouseEvent) => {
    e.preventDefault();
    try {

      const res = await fetch("http://localhost:8080/login" , {
        method : "POST",
        headers : {
          "Content-Type" : "application/json"
        },
        
        body : JSON.stringify({
          email : email ,
          password : password ,
        })
      })

      if(!res.ok){
        alert(await res.text());
        return ;
      }

      alert("login succesfully");

      const data = await res.json();

      localStorage.setItem("token", data.token);
      window.location.href = "/sessions";
      
    } catch (error) {
      console.log(error);
    }
  }
  
  return (
    <>
       <section className='mainlogin'>
        <div className='login'>
          <h1>login</h1>
          <input type="text" placeholder='email' className='logininp' onChange={(e) => setemail(e.target.value)}/>
          <input type="text" placeholder='password' className='logininp' onChange={(e) => setpassword(e.target.value)} />
          <input type="button" value = 'login' className='submitbtn' onClick={login}/>
            <h3>
              Don't have an account? <Link to="/signup">Sign up</Link>
            </h3>
       </div>
       </section>
    </>
  )
}

export default Login