import React, { useState } from 'react'

import './signup.css'
function Signup() {
 
const [name , setname] = useState("");
const [email , setemail] = useState("");
const [password , setpassword] = useState("");


const createuser = async ( e: React.MouseEvent) => {
  e.preventDefault() ;
  try {
    const res = await fetch
  } catch (error) {
    
  }
}
  
}
  return (
    <>
       <section>
        <div>
          <input type="text" placeholder='name' onChange={(e) => setname(e.target.value)} />
          <input type="text" placeholder='email' onChange={(e) => setemail(e.target.value)}/>
          <input type="text" placeholder='password' onChange={(e) => setpassword(e.target.value)}/>
          <input type="button" value= "create acount" onClick={createuser}/>
        </div>
       </section>
    </>
  )
}

export default Signup