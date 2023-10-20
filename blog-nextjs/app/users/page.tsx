import getUsers from "@/lib/getUsers";
import Link from "next/link";
import React, { Suspense } from "react";

const UsersPage = async () => {
  const users = await getUsers();
  const content = (
    <main>
      <h1>Users</h1>
      <Suspense fallback="<p>Loading...</p>">
        <ol type="1">
          {users.map((user) => {
            return (
              <li key={user.id}>
                <Link href={"/users/" + user.id}>{user.name}</Link>
              </li>
            );
          })}
        </ol>
      </Suspense>
    </main>
  );
  return content;
};

export default UsersPage;
