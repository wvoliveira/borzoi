import Link from "next/link";
import useUser from "../lib/iron/useUser";
import {useRouter} from "next/router";
import Image from "next/image";
import {AuthAPI} from "../lib/api/auth";

export default function Header() {
    const {user, mutateUser} = useUser();
    const router = useRouter();

    return (<header>
        <nav>
            <ul>
                <li>
                    <Link href="/">
                        <a>Home</a>
                    </Link>
                </li>
                {!user && (<>
                        <li>
                            <Link href="/login">
                                <a>Login</a>
                            </Link>
                        </li>
                        <li>
                            <Link href="/register">
                                <a>Register</a>
                            </Link>
                        </li>
                    </>
                )}
                {user?.status === "successful" && (<>
                    <li>
                        <Link href="/profile">
                            <a>
                    <span
                        style={{
                            marginRight: ".3em", verticalAlign: "middle", borderRadius: "100%", overflow: "hidden",
                        }}
                    >
                      {/*<Image*/}
                        {/*    src={user.avatarUrl}*/}
                        {/*    width={32}*/}
                        {/*    height={32}*/}
                        {/*    alt=""*/}
                        {/*/>*/}
                    </span>
                                Profile (Static Generation, recommended)
                            </a>
                        </Link>
                    </li>
                    <li>
                        {/* In this case, we're fine with linking with a regular a in case of no JavaScript */}
                        {/* eslint-disable-next-line @next/next/no-html-link-for-pages */}
                        <a
                            href="/"
                            onClick={async (e) => {
                                e.preventDefault();
                                await AuthAPI.Logout();
                                mutateUser(false);
                                await router.push("/login");
                            }}
                        >
                            Logout
                        </a>
                    </li>
                </>)}
                <li>
                    <a href="https://github.com/vvo/iron-session">
                        <Image
                            src="/GitHub-Mark-Light-32px.png"
                            width="32"
                            height="32"
                            alt=""
                        />
                    </a>
                </li>
            </ul>
        </nav>
        <style jsx>{`
        ul {
          display: flex;
          list-style: none;
          margin-left: 0;
          padding-left: 0;
        }
        li {
          margin-right: 1rem;
          display: flex;
        }
        li:first-child {
          margin-left: auto;
        }
        a {
          color: #fff;
          text-decoration: none;
          display: flex;
          align-items: center;
        }
        a img {
          margin-right: 1em;
        }
        header {
          padding: 0.2rem;
          color: #fff;
          background-color: #333;
        }
      `}</style>
    </header>);
}
