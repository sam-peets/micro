"use client";

import { useEffect, useState } from 'react';
import { ApiCall, GetRecent, GetUser } from './lib/api';
import { Post, User } from './types';
import PostList from './components/postList';
import CreatePost from './components/createPost';

// page content
export default function Page() {
  const [data, setData] = useState<Post[] | []>([]);

  useEffect(() => {
    GetRecent(10, 0).then(x => {
      setData(x)
    })
  }, [])
  return (
    <main>
      <CreatePost/>
      <PostList posts={data}></PostList>
    </main>
  );
}