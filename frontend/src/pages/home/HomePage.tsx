import { MainLayout } from '../../components/layout/MainLayout';
import { PostCard } from '../../components/blog/PostCard';
import { Button } from '../../components/ui/Button';
import { usePosts } from '../../hooks/usePosts';
import { Link } from 'react-router-dom';

export function HomePage() {
  const { posts, loading, error } = usePosts({ limit: 3 });

  return (
    <MainLayout>
      <section className="py-12">
        <div className="text-center mb-12">
          <h1 className="text-4xl font-extrabold text-gray-900 sm:text-5xl sm:tracking-tight">
            Welcome to MiniiBlog
          </h1>
          <p className="mt-5 max-w-xl mx-auto text-xl text-gray-500">
            A simple blog platform showcasing the latest thoughts and ideas.
          </p>
        </div>
      </section>

      <section className="py-8">
        <div className="flex items-center justify-between mb-8">
          <h2 className="text-2xl font-bold text-gray-900">Latest Posts</h2>
          <Link to="/blog">
            <Button variant="outline" size="sm">
              View all posts
            </Button>
          </Link>
        </div>

        {loading ? (
          <div className="flex justify-center items-center h-64">
            <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-blue-500"></div>
          </div>
        ) : error ? (
          <div className="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded">
            <p>Error loading posts. Please try again later.</p>
          </div>
        ) : posts.length === 0 ? (
          <div className="bg-gray-50 border border-gray-200 text-gray-700 px-4 py-5 rounded">
            <p>No posts found. Check back later for new content!</p>
          </div>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            {posts.map((post) => (
              <PostCard key={post.id} post={post} />
            ))}
          </div>
        )}
      </section>
    </MainLayout>
  );
}
