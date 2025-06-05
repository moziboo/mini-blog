import { useParams, Link } from 'react-router-dom';
import { usePost } from '../../hooks/usePosts';
import { Button } from '../../components/ui/Button';
import ReactMarkdown from 'react-markdown';
import remarkGfm from 'remark-gfm';

export function PostPage() {
  const { id } = useParams<{ id: string }>();
  const { post, loading, error } = usePost(id || '');

  if (loading) {
    return (
      <div className="flex justify-center items-center h-64">
        <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-blue-500"></div>
      </div>
    );
  }

  if (error || !post) {
    return (
      <div className="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded">
        <h2 className="text-xl font-bold mb-2">Error Loading Post</h2>
        <p>{error?.message || 'The requested post could not be found.'}</p>
        <div className="mt-4">
          <Link to="/">
            <Button variant="outline">Back to Feed</Button>
          </Link>
        </div>
      </div>
    );
  }

  const { title, content, author, publishedAt } = post;

  const formattedDate = publishedAt
    ? new Date(publishedAt).toLocaleDateString('en-US', {
        year: 'numeric',
        month: 'long',
        day: 'numeric',
      })
    : null;

  return (
    <article className="max-w-3xl mx-auto">
      <header className="mb-8">
        <h1 className="text-3xl sm:text-4xl font-bold text-gray-900 mb-4">
          {title}
        </h1>

        <div className="flex items-center text-gray-500">
          {author && <span className="font-medium">{author}</span>}
          {formattedDate && (
            <>
              <span className="mx-2">•</span>
              <time dateTime={publishedAt}>{formattedDate}</time>
            </>
          )}
        </div>
      </header>

      <div className="prose prose-lg max-w-none prose-headings:font-bold prose-h1:text-3xl prose-h2:text-2xl prose-h3:text-xl prose-h4:text-lg">
        <ReactMarkdown 
          remarkPlugins={[remarkGfm]}
        >
          {content}
        </ReactMarkdown>
      </div>
    </article>
  );
}
