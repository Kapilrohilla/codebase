import { addPost, getPosts } from "@/@lib/data";
import { NextResponse } from "next/server";

export const GET = async (req: Request, res: Response) => {
  try {
    const posts = getPosts();
    return NextResponse.json(posts, { status: 200 });
  } catch (err) {
    return NextResponse.json({ message: "Error", err }, { status: 500 });
  }
};

export const POST = async (req: Request, res: Response) => {
  // Get the request body as JSON
  const { title, desc } = await req.json();
  const post = {
    id: Date.now().toString(),
    title,
    desc,
    date: new Date(),
  };
  try {
    addPost(post);
    return NextResponse.json({ post }, { status: 200 });
  } catch (err) {
    return NextResponse.json({ message: "Error", err }, { status: 500 });
  }
};
