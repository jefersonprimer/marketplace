-- WARNING: This schema is for context only and is not meant to be run.
-- Table order and constraints may not be valid for execution.

Enums
order_status {
    pending,
    paid, 
    shipped, 
    delivered, 
    cancelled
}

user_role {
    admin, 
    seller, 
    customer
}

product_status {
    active, 
    inactive, 
    out_of_stock
}


CREATE TABLE public.addresses (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  user_id uuid,
  full_name text NOT NULL,
  phone text NOT NULL,
  street text NOT NULL,
  number text,
  complement text,
  neighborhood text,
  city text NOT NULL,
  state text NOT NULL,
  zip_code text NOT NULL,
  country text DEFAULT 'Brasil'::text,
  is_default boolean DEFAULT false,
  created_at timestamp with time zone DEFAULT now(),
  latitude numeric,
  longitude numeric,
  type text CHECK (type = ANY (ARRAY['shipping'::text, 'billing'::text, 'home'::text, 'work'::text])),
  store_id uuid,
  CONSTRAINT addresses_pkey PRIMARY KEY (id),
  CONSTRAINT addresses_store_id_fkey FOREIGN KEY (store_id) REFERENCES public.stores(id),
  CONSTRAINT addresses_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.profiles(id)
);
CREATE TABLE public.cart_items (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  user_id uuid NOT NULL,
  product_id uuid NOT NULL,
  quantity integer NOT NULL CHECK (quantity > 0),
  created_at timestamp with time zone DEFAULT now(),
  CONSTRAINT cart_items_pkey PRIMARY KEY (id),
  CONSTRAINT cart_items_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id),
  CONSTRAINT cart_items_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.profiles(id)
);
CREATE TABLE public.categories (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  name text NOT NULL,
  slug text NOT NULL UNIQUE,
  parent_id uuid,
  created_at timestamp with time zone DEFAULT now(),
  CONSTRAINT categories_pkey PRIMARY KEY (id),
  CONSTRAINT categories_parent_id_fkey FOREIGN KEY (parent_id) REFERENCES public.categories(id)
);
CREATE TABLE public.coupon_categories (
  coupon_id uuid NOT NULL,
  category_id uuid NOT NULL,
  CONSTRAINT coupon_categories_pkey PRIMARY KEY (coupon_id, category_id),
  CONSTRAINT coupon_categories_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.categories(id),
  CONSTRAINT coupon_categories_coupon_id_fkey FOREIGN KEY (coupon_id) REFERENCES public.coupons(id)
);
CREATE TABLE public.coupon_products (
  coupon_id uuid NOT NULL,
  product_id uuid NOT NULL,
  CONSTRAINT coupon_products_pkey PRIMARY KEY (coupon_id, product_id),
  CONSTRAINT coupon_products_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id),
  CONSTRAINT coupon_products_coupon_id_fkey FOREIGN KEY (coupon_id) REFERENCES public.coupons(id)
);
CREATE TABLE public.coupons (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  code text NOT NULL UNIQUE,
  discount_type text CHECK (discount_type = ANY (ARRAY['percentage'::text, 'fixed'::text])),
  discount_value numeric NOT NULL CHECK (discount_value >= 0::numeric),
  max_uses integer,
  used_count integer DEFAULT 0,
  valid_from timestamp with time zone,
  valid_until timestamp with time zone,
  created_at timestamp with time zone DEFAULT now(),
  store_id uuid,
  min_order_value numeric,
  applies_to text DEFAULT 'all'::text CHECK (applies_to = ANY (ARRAY['all'::text, 'products'::text, 'categories'::text, 'shipping'::text])),
  CONSTRAINT coupons_pkey PRIMARY KEY (id),
  CONSTRAINT coupons_store_id_fkey FOREIGN KEY (store_id) REFERENCES public.stores(id)
);
CREATE TABLE public.deliveries (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  order_id uuid,
  delivery_partner_id uuid,
  status text NOT NULL DEFAULT 'pending'::text,
  tracking_code text,
  estimated_delivery_at timestamp with time zone,
  delivered_at timestamp with time zone,
  created_at timestamp with time zone DEFAULT now(),
  CONSTRAINT deliveries_pkey PRIMARY KEY (id),
  CONSTRAINT deliveries_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.orders(id)
);
CREATE TABLE public.disputes (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  order_id uuid NOT NULL,
  opened_by uuid NOT NULL,
  reason text NOT NULL,
  status text DEFAULT 'open'::text CHECK (status = ANY (ARRAY['open'::text, 'in_review'::text, 'resolved'::text, 'rejected'::text])),
  resolution text,
  created_at timestamp with time zone DEFAULT now(),
  updated_at timestamp with time zone DEFAULT now(),
  CONSTRAINT disputes_pkey PRIMARY KEY (id),
  CONSTRAINT disputes_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.orders(id),
  CONSTRAINT disputes_opened_by_fkey FOREIGN KEY (opened_by) REFERENCES public.profiles(id)
);
CREATE TABLE public.marketplace_settings (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  key text NOT NULL UNIQUE,
  value text NOT NULL,
  updated_at timestamp with time zone DEFAULT now(),
  CONSTRAINT marketplace_settings_pkey PRIMARY KEY (id)
);
CREATE TABLE public.notifications (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  user_id uuid,
  type text CHECK (type = ANY (ARRAY['email'::text, 'whatsapp'::text, 'system'::text])),
  content text,
  status text CHECK (status = ANY (ARRAY['pending'::text, 'sent'::text, 'failed'::text])),
  created_at timestamp with time zone DEFAULT now(),
  CONSTRAINT notifications_pkey PRIMARY KEY (id),
  CONSTRAINT notifications_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.profiles(id)
);
CREATE TABLE public.order_groups (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  customer_id uuid,
  total_amount numeric NOT NULL,
  status text CHECK (status = ANY (ARRAY['pending'::text, 'paid'::text, 'cancelled'::text, 'refunded'::text])),
  payment_intent_id text,
  created_at timestamp with time zone DEFAULT now(),
  CONSTRAINT order_groups_pkey PRIMARY KEY (id),
  CONSTRAINT order_groups_customer_id_fkey FOREIGN KEY (customer_id) REFERENCES public.profiles(id)
);
CREATE TABLE public.order_items (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  order_id uuid,
  product_id uuid,
  quantity integer NOT NULL,
  price numeric NOT NULL,
  CONSTRAINT order_items_pkey PRIMARY KEY (id),
  CONSTRAINT order_items_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.orders(id),
  CONSTRAINT order_items_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id)
);
CREATE TABLE public.order_messages (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  order_id uuid,
  sender_role text CHECK (sender_role = ANY (ARRAY['customer'::text, 'seller'::text, 'system'::text])),
  message text NOT NULL,
  created_at timestamp with time zone DEFAULT now(),
  CONSTRAINT order_messages_pkey PRIMARY KEY (id),
  CONSTRAINT order_messages_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.orders(id)
);
CREATE TABLE public.order_status_history (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  order_id uuid,
  status text CHECK (status = ANY (ARRAY['pending'::text, 'paid'::text, 'shipped'::text, 'delivered'::text, 'cancelled'::text])),
  changed_at timestamp with time zone DEFAULT now(),
  CONSTRAINT order_status_history_pkey PRIMARY KEY (id),
  CONSTRAINT order_status_history_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.orders(id)
);
CREATE TABLE public.orders (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  store_id uuid,
  total_amount numeric NOT NULL,
  status USER-DEFINED DEFAULT 'pending'::order_status,
  created_at timestamp with time zone DEFAULT now(),
  shipping_address_id uuid,
  transaction_id uuid,
  order_group_id uuid,
  updated_at timestamp with time zone DEFAULT now(),
  commission_fee numeric DEFAULT 0,
  net_amount numeric DEFAULT (total_amount - commission_fee),
  CONSTRAINT orders_pkey PRIMARY KEY (id),
  CONSTRAINT orders_store_id_fkey FOREIGN KEY (store_id) REFERENCES public.stores(id),
  CONSTRAINT orders_shipping_address_id_fkey FOREIGN KEY (shipping_address_id) REFERENCES public.addresses(id),
  CONSTRAINT orders_payments_id_fkey FOREIGN KEY (order_group_id) REFERENCES public.order_groups(id)
);
CREATE TABLE public.payment_attempts (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  payment_id uuid NOT NULL,
  attempt_number integer,
  status text CHECK (status = ANY (ARRAY['pending'::text, 'success'::text, 'failed'::text])),
  raw_response jsonb,
  created_at timestamp with time zone DEFAULT now(),
  CONSTRAINT payment_attempts_pkey PRIMARY KEY (id),
  CONSTRAINT payment_attempts_payment_id_fkey FOREIGN KEY (payment_id) REFERENCES public.payments(id)
);
CREATE TABLE public.payments (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  order_group_id uuid,
  gateway text,
  gateway_payment_id text,
  amount numeric,
  status text CHECK (status = ANY (ARRAY['pending'::text, 'paid'::text, 'failed'::text, 'refunded'::text])),
  raw_response jsonb,
  created_at timestamp with time zone DEFAULT now(),
  CONSTRAINT payments_pkey PRIMARY KEY (id),
  CONSTRAINT group_payments_order_group_id_fkey FOREIGN KEY (order_group_id) REFERENCES public.order_groups(id)
);
CREATE TABLE public.product_attributes (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  product_id uuid NOT NULL,
  name text NOT NULL,
  value text NOT NULL,
  created_at timestamp with time zone DEFAULT now(),
  CONSTRAINT product_attributes_pkey PRIMARY KEY (id),
  CONSTRAINT product_attributes_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id)
);
CREATE TABLE public.product_images (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  product_id uuid,
  image_url text NOT NULL,
  is_main boolean DEFAULT false,
  created_at timestamp with time zone DEFAULT now(),
  CONSTRAINT product_images_pkey PRIMARY KEY (id),
  CONSTRAINT product_images_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id)
);
CREATE TABLE public.product_price_history (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  product_id uuid NOT NULL,
  old_price numeric NOT NULL,
  new_price numeric NOT NULL,
  changed_at timestamp with time zone DEFAULT now(),
  CONSTRAINT product_price_history_pkey PRIMARY KEY (id),
  CONSTRAINT product_price_history_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id)
);
CREATE TABLE public.product_reviews (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  product_id uuid,
  user_id uuid,
  rating integer CHECK (rating >= 1 AND rating <= 5),
  comment text,
  created_at timestamp with time zone DEFAULT now(),
  CONSTRAINT product_reviews_pkey PRIMARY KEY (id),
  CONSTRAINT product_reviews_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id),
  CONSTRAINT product_reviews_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.profiles(id)
);
CREATE TABLE public.product_variants (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  product_id uuid NOT NULL,
  name text NOT NULL,
  CONSTRAINT product_variants_pkey PRIMARY KEY (id),
  CONSTRAINT product_variants_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id)
);
CREATE TABLE public.products (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  store_id uuid NOT NULL,
  name text NOT NULL,
  slug text NOT NULL UNIQUE,
  description text,
  price numeric NOT NULL,
  stock integer DEFAULT 0,
  status USER-DEFINED DEFAULT 'active'::product_status,
  image_url text,
  created_at timestamp with time zone DEFAULT now(),
  category_id uuid NOT NULL,
  updated_at timestamp with time zone DEFAULT now(),
  sku text UNIQUE,
  gtin text,
  weight numeric,
  length numeric,
  width numeric,
  height numeric,
  CONSTRAINT products_pkey PRIMARY KEY (id),
  CONSTRAINT products_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.categories(id),
  CONSTRAINT products_store_id_fkey FOREIGN KEY (store_id) REFERENCES public.stores(id)
);
CREATE TABLE public.profile_documents (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  profile_id uuid NOT NULL,
  document_type text CHECK (document_type = ANY (ARRAY['cpf'::text, 'cnpj'::text, 'rg'::text, 'cnh'::text, 'passport'::text, 'address_proof'::text])),
  document_number text,
  document_url text,
  status text DEFAULT 'pending'::text CHECK (status = ANY (ARRAY['pending'::text, 'approved'::text, 'rejected'::text])),
  created_at timestamp with time zone DEFAULT now(),
  CONSTRAINT profile_documents_pkey PRIMARY KEY (id),
  CONSTRAINT profile_documents_profile_id_fkey FOREIGN KEY (profile_id) REFERENCES public.profiles(id)
);
CREATE TABLE public.profiles (
  id uuid NOT NULL,
  full_name text,
  role USER-DEFINED NOT NULL DEFAULT 'customer'::user_role,
  avatar_url text,
  created_at timestamp with time zone DEFAULT now(),
  cpf text UNIQUE,
  cnpj text UNIQUE,
  phone text,
  is_seller boolean DEFAULT false,
  updated_at timestamp with time zone DEFAULT now(),
  CONSTRAINT profiles_pkey PRIMARY KEY (id),
  CONSTRAINT profiles_id_fkey FOREIGN KEY (id) REFERENCES auth.users(id)
);
CREATE TABLE public.promotion_categories (
  promotion_id uuid NOT NULL,
  category_id uuid NOT NULL,
  CONSTRAINT promotion_categories_pkey PRIMARY KEY (promotion_id, category_id),
  CONSTRAINT promotion_categories_promotion_id_fkey FOREIGN KEY (promotion_id) REFERENCES public.promotions(id),
  CONSTRAINT promotion_categories_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.categories(id)
);
CREATE TABLE public.promotion_products (
  promotion_id uuid NOT NULL,
  product_id uuid NOT NULL,
  CONSTRAINT promotion_products_pkey PRIMARY KEY (promotion_id, product_id),
  CONSTRAINT promotion_products_promotion_id_fkey FOREIGN KEY (promotion_id) REFERENCES public.promotions(id),
  CONSTRAINT promotion_products_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id)
);
CREATE TABLE public.promotions (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  store_id uuid NOT NULL,
  name text NOT NULL,
  description text,
  discount_type text NOT NULL CHECK (discount_type = ANY (ARRAY['percentage'::text, 'fixed'::text])),
  discount_value numeric NOT NULL,
  start_date timestamp with time zone NOT NULL,
  end_date timestamp with time zone NOT NULL,
  applies_to text DEFAULT 'all'::text CHECK (applies_to = ANY (ARRAY['all'::text, 'products'::text, 'categories'::text])),
  active boolean DEFAULT true,
  created_at timestamp with time zone DEFAULT now(),
  CONSTRAINT promotions_pkey PRIMARY KEY (id),
  CONSTRAINT promotions_store_id_fkey FOREIGN KEY (store_id) REFERENCES public.stores(id)
);
CREATE TABLE public.returns (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  order_id uuid NOT NULL,
  order_item_id uuid NOT NULL,
  reason text NOT NULL,
  status text DEFAULT 'requested'::text CHECK (status = ANY (ARRAY['requested'::text, 'approved'::text, 'rejected'::text, 'returned'::text, 'refunded'::text])),
  requested_at timestamp with time zone DEFAULT now(),
  processed_at timestamp with time zone,
  CONSTRAINT returns_pkey PRIMARY KEY (id),
  CONSTRAINT returns_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.orders(id),
  CONSTRAINT returns_order_item_id_fkey FOREIGN KEY (order_item_id) REFERENCES public.order_items(id)
);
CREATE TABLE public.seller_payments (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  order_id uuid NOT NULL,
  store_id uuid NOT NULL,
  amount numeric NOT NULL,
  status text DEFAULT 'pending'::text CHECK (status = ANY (ARRAY['pending'::text, 'processing'::text, 'paid'::text, 'failed'::text])),
  payout_date timestamp with time zone,
  created_at timestamp with time zone DEFAULT now(),
  CONSTRAINT seller_payments_pkey PRIMARY KEY (id),
  CONSTRAINT seller_payments_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.orders(id),
  CONSTRAINT seller_payments_store_id_fkey FOREIGN KEY (store_id) REFERENCES public.stores(id)
);
CREATE TABLE public.shipping_methods (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  name text NOT NULL,
  description text,
  cost numeric DEFAULT 0,
  estimated_days integer,
  created_at timestamp with time zone DEFAULT now(),
  CONSTRAINT shipping_methods_pkey PRIMARY KEY (id)
);
CREATE TABLE public.stock_movements (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  product_id uuid,
  change integer NOT NULL,
  reason text,
  created_at timestamp with time zone DEFAULT now(),
  CONSTRAINT stock_movements_pkey PRIMARY KEY (id),
  CONSTRAINT stock_movements_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id)
);
CREATE TABLE public.store_categories (
  store_id uuid NOT NULL,
  category_id uuid NOT NULL,
  CONSTRAINT store_categories_pkey PRIMARY KEY (store_id, category_id),
  CONSTRAINT store_categories_store_id_fkey FOREIGN KEY (store_id) REFERENCES public.stores(id),
  CONSTRAINT store_categories_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.categories(id)
);
CREATE TABLE public.store_commissions (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  store_id uuid NOT NULL,
  commission_type text DEFAULT 'percentage'::text CHECK (commission_type = ANY (ARRAY['percentage'::text, 'fixed'::text])),
  commission_value numeric NOT NULL,
  valid_from timestamp with time zone DEFAULT now(),
  valid_to timestamp with time zone,
  CONSTRAINT store_commissions_pkey PRIMARY KEY (id),
  CONSTRAINT store_commissions_store_id_fkey FOREIGN KEY (store_id) REFERENCES public.stores(id)
);
CREATE TABLE public.store_policies (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  store_id uuid NOT NULL,
  policy_type text CHECK (policy_type = ANY (ARRAY['return'::text, 'shipping'::text, 'privacy'::text, 'terms'::text])),
  content text NOT NULL,
  created_at timestamp with time zone DEFAULT now(),
  CONSTRAINT store_policies_pkey PRIMARY KEY (id),
  CONSTRAINT store_policies_store_id_fkey FOREIGN KEY (store_id) REFERENCES public.stores(id)
);
CREATE TABLE public.store_reviews (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  store_id uuid,
  user_id uuid,
  rating integer CHECK (rating >= 1 AND rating <= 5),
  comment text,
  created_at timestamp with time zone DEFAULT now(),
  CONSTRAINT store_reviews_pkey PRIMARY KEY (id),
  CONSTRAINT store_reviews_store_id_fkey FOREIGN KEY (store_id) REFERENCES public.stores(id),
  CONSTRAINT store_reviews_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.profiles(id)
);
CREATE TABLE public.store_shipping_methods (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  store_id uuid NOT NULL,
  shipping_method_id uuid NOT NULL,
  active boolean DEFAULT true,
  created_at timestamp with time zone DEFAULT now(),
  CONSTRAINT store_shipping_methods_pkey PRIMARY KEY (id),
  CONSTRAINT store_shipping_methods_shipping_method_id_fkey FOREIGN KEY (shipping_method_id) REFERENCES public.shipping_methods(id),
  CONSTRAINT store_shipping_methods_store_id_fkey FOREIGN KEY (store_id) REFERENCES public.stores(id)
);
CREATE TABLE public.store_transactions (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  store_id uuid NOT NULL,
  type text NOT NULL CHECK (type = ANY (ARRAY['credit'::text, 'debit'::text])),
  amount numeric NOT NULL,
  description text,
  related_order_id uuid,
  created_at timestamp with time zone DEFAULT now(),
  CONSTRAINT store_transactions_pkey PRIMARY KEY (id),
  CONSTRAINT store_transactions_store_id_fkey FOREIGN KEY (store_id) REFERENCES public.stores(id),
  CONSTRAINT store_transactions_related_order_id_fkey FOREIGN KEY (related_order_id) REFERENCES public.orders(id)
);
CREATE TABLE public.stores (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  owner_id uuid,
  name text NOT NULL,
  slug text NOT NULL UNIQUE,
  description text,
  logo_url text,
  banner_url text,
  created_at timestamp with time zone DEFAULT now(),
  updated_at timestamp with time zone DEFAULT now(),
  CONSTRAINT stores_pkey PRIMARY KEY (id),
  CONSTRAINT stores_owner_id_fkey FOREIGN KEY (owner_id) REFERENCES public.profiles(id)
);
CREATE TABLE public.user_activity_log (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  user_id uuid NOT NULL,
  action text NOT NULL,
  metadata jsonb,
  created_at timestamp with time zone DEFAULT now(),
  CONSTRAINT user_activity_log_pkey PRIMARY KEY (id),
  CONSTRAINT user_activity_log_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.profiles(id)
);
CREATE TABLE public.variant_options (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  variant_id uuid NOT NULL,
  value text NOT NULL,
  price_adjustment numeric DEFAULT 0,
  stock integer NOT NULL DEFAULT 0 CHECK (stock >= 0),
  created_at timestamp with time zone DEFAULT now(),
  CONSTRAINT variant_options_pkey PRIMARY KEY (id),
  CONSTRAINT variant_options_variant_id_fkey FOREIGN KEY (variant_id) REFERENCES public.product_variants(id)
);
CREATE TABLE public.warehouses (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  store_id uuid NOT NULL,
  name text NOT NULL,
  address_id uuid,
  created_at timestamp with time zone DEFAULT now(),
  CONSTRAINT warehouses_pkey PRIMARY KEY (id),
  CONSTRAINT warehouses_store_id_fkey FOREIGN KEY (store_id) REFERENCES public.stores(id),
  CONSTRAINT warehouses_address_id_fkey FOREIGN KEY (address_id) REFERENCES public.addresses(id)
);
CREATE TABLE public.wishlists (
  id uuid NOT NULL DEFAULT gen_random_uuid(),
  user_id uuid NOT NULL,
  product_id uuid NOT NULL,
  created_at timestamp with time zone DEFAULT now(),
  CONSTRAINT wishlists_pkey PRIMARY KEY (id),
  CONSTRAINT wishlists_product_id_fkey FOREIGN KEY (product_id) REFERENCES public.products(id),
  CONSTRAINT wishlists_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.profiles(id)
);