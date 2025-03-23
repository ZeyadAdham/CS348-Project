import { useState, useEffect } from 'react';

function App() {
    const [message, setMessage] = useState('Loading...');

    useEffect(() => {
        const fetchMessage = async () => {
            try {
                const response = await fetch('http://localhost:4000/api/hello');
                if (!response.ok) {
                    throw new Error(`HTTP error! status: ${response.status}`);
                }
                const data = await response.json();
                setMessage(data.message);
            } catch (error) {
                console.error('Error fetching data:', error);
                setMessage('Failed to fetch message from backend');
            }
        };
        fetchMessage();
    }, []);

    return (
        <div>
            <h2>Hello, World! (from Frontend)</h2>
            <h2>{message}</h2>
        </div>
    );
}

export default App;
