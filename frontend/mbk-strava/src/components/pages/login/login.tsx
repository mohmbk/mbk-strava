import { useState } from 'react'
import { Link } from 'react-router-dom'

import './login.css'
function Login() {
 
  const [email , setemail] = useState('');
  const [ password , setpassword] = useState('');

  
  return (
    <>
       <section className='mainlogin'>
        <div className='login'>
          <h1>login</h1>
          <input type="text" placeholder='email' className='logininp' onChange={(e) => setemail(e.target.value)}/>
          <input type="text" placeholder='password' className='logininp' onChange={(e) => setpassword(e.target.value)} />
          <input type="button" value = 'login' className='submitbtn'/>
            <h3>
              Don't have an account? <Link to="/signup">Sign up</Link>
            </h3>
       </div>
       </section>
    </>
  )
}

export default Login