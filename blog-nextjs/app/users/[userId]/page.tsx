import getUserBlogs from "@/lib/getUserBlogs";
import PostTitle from "./component/Post";
import { Metadata, ResolvingMetadata } from "next";
import getUser from "@/lib/getUser";

type Props = {
  params: { userId: string };
  searchParams: { [key: string]: string | string[] | undefined };
};

export async function generateMetadata(
  { params, searchParams }: Props,
  parent: ResolvingMetadata
): Promise<Metadata> {
  // read route params
  const id = params.userId;

  // fetch data
  const user = await getUser(id);
  return {
    title: user.name + " blogs",
  };
}

const UserPost = async ({ params }: { params: { userId: string } }) => {
  const userBlogs = await getUserBlogs(params.userId);
  return (
    <section>
      <h1>This is userSpecific blogs where userId = {params.userId}</h1>
      <ol>
        {userBlogs.map((blog) => {
          return <PostTitle post={blog} key={blog.id} />;
        })}
      </ol>
    </section>
  );
};

export default UserPost;
