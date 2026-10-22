UPDATE products
SET 
    title = 'Updated Mechanical Gaming Keyboard Pro',
    description = 'Upgraded wireless RGB mechanical keyboard with hot-swappable red switches and 2.4GHz receiver.',
    price = 99.99,
    img_url = 'https://example.com/images/keyboard-pro.jpg'
WHERE id = 2
RETURNING id, title, price, img_url;