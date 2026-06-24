import { useEffect, useState } from 'react'

import './statistics.css'
function Statistics() {
 
  interface stat {
    totaldistance: number,
    totaltime: number,
    averagetotalpace: number,
  }

  const [stat, setstat] = useState<stat | null>(null);

  useEffect(() =>{ 

      async function fetchstat() {
        const token = localStorage.getItem("token");
        const response = await fetch("http://localhost:8080/sessions" , {
          method : "GET",
          headers : {
            "Authorization": `Bearer ${token}`
          }
        })

        const data = await response.json();
        setstat(data);
      }

      fetchstat();



  } , [] )

  return (
    <>
       <div>
        <h3>{stat?.totaldistance}</h3>
        <h3>{stat?.totaltime}</h3>
        <h3>{stat?.averagetotalpace}</h3>
       </div>
    </>
  )
}

export default Statistics