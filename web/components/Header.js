import useUser from "../lib/utils/useUser";
import {useRouter} from "next/router";
import {AuthAPI} from "../lib/api/auth";
import * as React from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import HomeIcon from '@mui/icons-material/Home';
import Container from '@mui/material/Container';
import {Link} from "@mui/material";

export default function Header() {
    const {user, mutateUser} = useUser();
    const router = useRouter();

    return (<header>
        <Box sx={{flexGrow: 1}}>
            <AppBar position="static">
                <Container maxWidth="xl">
                <Toolbar disableGutters>
                    <Typography variant="h6" noWrap component="div" sx={{flexGrow: 1, display: { xs: 'flex'}}}>
                            <a onClick={() => {router.push("/")}}>BORZOI</a>
                    </Typography>
                    {!user && (<>
                        <Button onClick={() => {router.push("/login")}} color="inherit">Login</Button>
                        <Button onClick={() => {router.push("/register")}} color="inherit">Register</Button>
                    </>)}
                    {user?.status === "successful" && (<>
                        <Button onClick={() => {router.push("/profile")}} color="inherit">Profile</Button>
                        <Button onClick={async (e) => {
                            e.preventDefault();
                            await AuthAPI.Logout();
                            mutateUser(false);
                            await router.push("/");
                        }} color="inherit">Logout</Button>
                    </>)}
                </Toolbar>
                </Container>
            </AppBar>
        </Box>
    </header>)
    //     <header>
    //     <nav>
    //         <ul>
    //             <li>
    //                 <Link href="/">
    //                     <a>Home</a>
    //                 </Link>
    //             </li>
    //             {!user && (<>
    //                     <li>
    //                         <Link href="/login">
    //                             <a>Login</a>
    //                         </Link>
    //                     </li>
    //                     <li>
    //                         <Link href="/register">
    //                             <a>Register</a>
    //                         </Link>
    //                     </li>
    //                 </>
    //             )}
    //             {user?.status === "successful" && (<>
    //                 <li>
    //                     <Link href="/profile">
    //                         <a>
    //                 <span
    //                     style={{
    //                         marginRight: ".3em", verticalAlign: "middle", borderRadius: "100%", overflow: "hidden",
    //                     }}
    //                 >
    //                 </span>
    //                             Profile (Static Generation, recommended)
    //                         </a>
    //                     </Link>
    //                 </li>
    //                 <li>
    //                     {/* In this case, we're fine with linking with a regular a in case of no JavaScript */}
    //                     {/* eslint-disable-next-line @next/next/no-html-link-for-pages */}
    //                     <a
    //                         href="/"
    //                         onClick={async (e) => {
    //                             e.preventDefault();
    //                             await AuthAPI.Logout();
    //                             mutateUser(false);
    //                             await router.push("/login");
    //                         }}
    //                     >
    //                         Logout
    //                     </a>
    //                 </li>
    //             </>)}
    //         </ul>
    //     </nav>
    //     <style jsx>{`
    //     ul {
    //       display: flex;
    //       list-style: none;
    //       margin-left: 0;
    //       padding-left: 0;
    //     }
    //     li {
    //       margin-right: 1rem;
    //       display: flex;
    //     }
    //     li:first-child {
    //       margin-left: auto;
    //     }
    //     a {
    //       color: #fff;
    //       text-decoration: none;
    //       display: flex;
    //       align-items: center;
    //     }
    //     a img {
    //       margin-right: 1em;
    //     }
    //     header {
    //       padding: 0.2rem;
    //       color: #fff;
    //       background-color: #333;
    //     }
    //   `}</style>
    // </header>);
}
