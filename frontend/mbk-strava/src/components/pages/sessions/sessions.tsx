import { useEffect, useState } from 'react'

import './sessions.css'
function Sessions() {
 
interface Session {
    id : number ,
    title : string ,
    distance : number ,
    time : number ,
  }

  const [sessions , setsessions] = useState<Session[]>([]) ;

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
    </>
  )
}

export default Sessions