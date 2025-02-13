'use client'
 
import { useSearchParams } from 'next/navigation'
import { GetUser, GetUserPosts } from '../lib/api'
import { Post, User } from '../types'
import { useEffect, useState } from 'react'
import PostList from '../components/postList'
 
export default function Page() {
  const [user, setUser] = useState<User | null>(null)
  const [posts, setPosts] = useState<Post[] | []>([])

  const searchParams = useSearchParams()
  const uid = searchParams.get('id')

  useEffect(() => {
    if (uid == null) {
      return;
    }
    GetUser(parseInt(uid)).then(u => {
      setUser(u)
    })
    GetUserPosts(parseInt(uid), 10, 0).then(p => {
      setPosts(p)
    })
  }, [uid])
  
  return (
    <main>
      <h1>{user?.username}</h1>
      <PostList posts={posts}/>
    </main>
  )
}