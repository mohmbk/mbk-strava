import React, { useEffect, useState } from 'react'

import './sessions.css'
function Sessions() {
 
interface Session {
    id : number ,
    title : string ,
    distance : number ,
    time : number ,
  }

  const [sessions , setsessions] = useState<Session[]>([]) ;
  const [title , settitle] = useState("");
  const [distance , setdistance] = useState("");
  const [time , settime] = useState("") ;

  useEffect(() => {

        async function fetchSessions() {
          const token = localStorage.getItem("token");
          const response = await fetch("http://localhost:8080/sessions" , {
            method : "GET",
            headers : ({
              "Authorization": `Bearer ${token}`
            })
          })

          const data = await response.json();
          setsessions(data) ;
        }

        fetchSessions();

    }, []);

    
    const createactivity = async (e : React.MouseEvent) => {
      e.preventDefault();
      try {
        const token = localStorage.getItem("token");
        const response = await fetch("http://localhost:8080/sessions" , {
          method : "POST",
          headers : {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`
          },

          body : JSON.stringify({
            title : title ,
            distance : Number(distance),
            time : Number(time),
          }),
        });

        if(!response.ok){
          alert(await response.text()) ;
          return ;
        }

        alert("activity created");
        setdistance("");
        settime("");
        settitle("");
      } catch (error) {
        console.log(error);
      }
    }

  return (
    <>
       <div>
        {sessions.map(Session => (
            <div key={Session.id}>
                 <h3> title : {Session.title}</h3>
                 <h3> distance : {Session.distance}</h3>
                 <h3>allure : {Session.time / Session.distance}</h3>
            </div>
        ))}
       </div>


       <section className='createsection'>
          <div className='creatediv'>
            <h1>create activity</h1>
            <input type="text" placeholder='tittle' className='sessioninp' onChange={(e) => settitle(e.target.value)} />
            <input type="text" placeholder='distance' className='sessioninp' onChange={(e) => setdistance(e.target.value)}/>
            <input type="text" placeholder='time' className='sessioninp' onChange={(e) => settime(e.target.value)}/>
            <input type="button"  value= 'add activity !' className='sessionbtn' onClick={createactivity} />
          </div>
       </section>
    </>
  )
}

export default Sessions