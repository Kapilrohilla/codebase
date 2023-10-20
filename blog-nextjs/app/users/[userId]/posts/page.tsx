import getUserBlogs from "@/lib/getUserBlogs";
import PostTitle from "./component/Post";

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
