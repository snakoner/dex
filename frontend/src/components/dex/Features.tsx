import { Card } from "@/components/ui/card";
import { ArrowLeftRight, Shield, Zap } from "lucide-react";

const Features = () => {
  const features = [
    {
      icon: <ArrowLeftRight className="w-12 h-12 text-purple-500" />,
      title: "Instant Swaps",
      description:
        "Swap tokens instantly with minimal slippage using automated market makers and deep liquidity pools",
      image:
        "https://images.unsplash.com/photo-1605792657660-596af9009e82?w=800&auto=format&fit=crop&q=60&ixlib=rb-4.0.3",
      gradientClass: "from-purple-500/20 via-transparent to-transparent",
      particleColor: "bg-purple-500",
    },
    {
      icon: <Shield className="w-12 h-12 text-blue-500" />,
      title: "Secure Trading",
      description:
        "Trade directly from your wallet with non-custodial, audited smart contracts for maximum security",
      image:
        "https://images.unsplash.com/photo-1639762681485-074b7f938ba0?w=800&auto=format&fit=crop&q=60&ixlib=rb-4.0.3",
      gradientClass: "from-blue-500/20 via-transparent to-transparent",
      particleColor: "bg-blue-500",
    },
    {
      icon: <Zap className="w-12 h-12 text-green-500" />,
      title: "Low Fees",
      description:
        "Benefit from competitive fees and efficient routing to get the best prices across multiple liquidity pools",
      image:
        "https://images.unsplash.com/photo-1518186285589-2f7649de83e0?w=800&auto=format&fit=crop&q=60&ixlib=rb-4.0.3",
      gradientClass: "from-green-500/20 via-transparent to-transparent",
      particleColor: "bg-green-500",
    },
  ];

  return (
    <div className="w-full max-w-6xl mx-auto px-4 pb-16">
      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
        {features.map((feature, index) => (
          <Card
            key={index}
            className="group overflow-hidden hover:shadow-xl transition-all duration-300 animate-slide-in relative"
            style={{ animationDelay: `${index * 150}ms` }}
          >
            {/* Animated background particles */}
            <div className="absolute inset-0 overflow-hidden">
              <div
                className={`absolute inset-0 bg-gradient-to-b ${feature.gradientClass} opacity-30`}
              />
              <div
                className={`absolute w-2 h-2 ${feature.particleColor} rounded-full animate-float`}
                style={{
                  left: "20%",
                  top: "30%",
                  animationDelay: "0s",
                  opacity: 0.3,
                }}
              />
              <div
                className={`absolute w-2 h-2 ${feature.particleColor} rounded-full animate-float`}
                style={{
                  left: "80%",
                  top: "60%",
                  animationDelay: "1s",
                  opacity: 0.3,
                }}
              />
              <div
                className={`absolute w-3 h-3 ${feature.particleColor} rounded-full animate-float`}
                style={{
                  left: "50%",
                  top: "45%",
                  animationDelay: "2s",
                  opacity: 0.2,
                }}
              />
            </div>

            <div className="relative h-48 overflow-hidden">
              <img
                src={feature.image}
                alt={feature.title}
                className="w-full h-full object-cover transition-transform duration-300 group-hover:scale-110"
              />
              <div className="absolute inset-0 bg-gradient-to-t from-background/90 to-background/20 flex items-center justify-center">
                {feature.icon}
              </div>
            </div>
            <div className="p-6 space-y-2 relative">
              <h3 className="text-xl font-semibold">{feature.title}</h3>
              <p className="text-muted-foreground">{feature.description}</p>
            </div>
          </Card>
        ))}
      </div>
    </div>
  );
};

export default Features;
