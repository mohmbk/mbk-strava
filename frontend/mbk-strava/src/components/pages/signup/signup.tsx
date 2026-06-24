import React, { useState } from 'react'

import './signup.css'
function Signup() {
 
const [name , setname] = useState("");
const [email , setemail] = useState("");
const [password , setpassword] = useState("");

interface user {
  name : string ;
  email : string ;
  password : String ;
}

const createuser = async ( e: React.MouseEvent) => {
  e.preventDefault() ;
  try {
    const response = await fetch("http://localhost:8080/signup" , {
      method : "POST",
      headers : {
        "Content-Type" : "application/json" ,
      },
      body : JSON.stringify({
          name : name,
          email : email,
          password : password,
      }),
    })

    if(!response.ok){
      console.log(await response.text());
      return ;
    }

    setname("");
    setemail("");
    setpassword("");
  
    alert("user created");
    window.location.href = "/login";
  } catch (error) {
    console.log(error);
  }
}
  

  return (
    <>
       <section className='signup'>
        <div className='signupdiv'>
          <h1>create acount</h1>
          <input type="text" placeholder='name' className='logininp' onChange={(e) => setname(e.target.value)} />
          <input type="email" placeholder='email' className='logininp' onChange={(e) => setemail(e.target.value)}/>
          <input type="text" placeholder='password' className='logininp' onChange={(e) => setpassword(e.target.value)}/>
          <input type="button" value= "create acount" className='submitbtn' onClick={createuser}/>
        </div>
       </section>
    </>
  )

}

export default Signup