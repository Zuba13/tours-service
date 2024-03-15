-- Insert 1
INSERT INTO public.equipment (id, name, description)
VALUES (1, 'Tent', 'A portable shelter typically used for camping');

-- Insert 2
INSERT INTO public.equipment (id, name, description)
VALUES (2, 'Backpack', 'A bag carried on the back and used to carry items while hiking or traveling');

-- Insert 3
INSERT INTO public.equipment (id, name, description)
VALUES (3, 'Sleeping bag', 'A warm padded bag used for sleeping outdoors, especially when camping');

-- Insert 4
INSERT INTO public.equipment (id, name, description)
VALUES (4, 'Cooking stove', 'A portable stove used for cooking meals while camping or backpacking');

-- Insert 5
INSERT INTO public.equipment (id, name, description)
VALUES (5, 'Hiking boots', 'Sturdy footwear designed for hiking and walking on rough terrain');


-- Insert 1
INSERT INTO public.tours (id, author_id, name, description, tags, difficult, price, status, publish_time, archive_time, distance, travel_time_and_method)
VALUES (1, 102, 'Mountain Adventure', 'An exciting adventure through the mountains', 'adventure, mountain', 3, 50.00, 1, NOW(), NOW() + INTERVAL '1 month', 50, '[]'::jsonb);

-- Insert 2
INSERT INTO public.tours (id, author_id, name, description, tags, difficult, price, status, publish_time, archive_time, distance, travel_time_and_method)
VALUES (2, 102, 'Coastal Exploration', 'Discover hidden gems along the coastline', 'adventure, nature', 2, 40.00, 1, NOW(), NOW() + INTERVAL '2 weeks', 30, '[]'::jsonb);

-- Insert 3
INSERT INTO public.tours (id, author_id, name, description, tags, difficult, price, status, publish_time, archive_time, distance, travel_time_and_method)
VALUES (3, 102, 'Historical City Tour', 'Explore the rich history of our city', 'history, culture', 1, 20.00, 1, NOW(), NOW() + INTERVAL '3 months', 10, '[]'::jsonb);

-- Insert 4
INSERT INTO public.tours (id, author_id, name, description, tags, difficult, price, status, publish_time, archive_time, distance, travel_time_and_method)
VALUES (4, 102, 'Forest Hiking Adventure', 'Hike through dense forests and scenic trails', 'adventure, nature, hiking', 3, 35.00, 1, NOW(), NOW() + INTERVAL '1 month', 40, '[]'::jsonb);

-- Insert 5
INSERT INTO public.tours (id, author_id, name, description, tags, difficult, price, status, publish_time, archive_time, distance, travel_time_and_method)
VALUES (5, 102, 'Urban Exploration', 'Discover hidden spots and urban landmarks', 'urban, adventure', 2, 30.00, 1, NOW(), NOW() + INTERVAL '2 weeks', 20, '[]'::jsonb);
