import Link from 'next/link'
import { useRouter } from 'next/router'
import React from 'react';
import { useEffect } from "react";

import ClientCreate from './ClientCreate';
import ClientList from './ClientList';


export default function ClientSection() {
  const router = useRouter()
  const [section, setSection] = React.useState("");

  useEffect(() => {
    if (!router.isReady) return;

    const hash = router.asPath.split('#')[1];
    if (hash != undefined) {
      setSection(hash);
    } else {
      setSection("create");
    }

  }, [router.isReady])

  const handleSection = (event) => {
    var id = event.target.id
    if (section != id) {
      setSection(id);
    }
  }

  return (
    <>
      <Link href="#create"><a id="create" onClick={handleSection}>Create</a></Link>
      {" "} | {" "} 
      <Link href="#list"><a id="list" onClick={handleSection}>List</a></Link>

      <br />
      <br />

      {section == "create" && <ClientCreate />}
      {section == "list" && <ClientList />}
    </>
  )
}